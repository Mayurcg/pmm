// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm/admin/cli/flags"
	"github.com/percona/pmm/admin/commands"
)

// UpgradeCommand is used by Kong for CLI flags and commands.
type UpgradeCommand struct {
	DockerImage            string `default:"percona/pmm-server:2" help:"Docker image to use to upgrade PMM Server. Defaults to latest version"`
	ContainerID            string `default:"pmm-server" help:"Container ID of the PMM Server to upgrade"`
	NewContainerName       string `help:"Name of the new container for PMM Server. If this flag is set, --new-container-name-prefix is ignored. Must be different from the current container name"` //nolint:lll
	NewContainerNamePrefix string `default:"pmm-server" help:"Prefix for the name of the new container for PMM Server"`
	DisableBackup          bool   `help:"Disable backup of data from Docker volumes before a PMM Server upgrade"`

	AssumeYes               bool `name:"yes" short:"y" help:"Assume yes for all prompts"`
	docker                  containerManager
	l                       *logrus.Entry
	waitBeforeContainerStop time.Duration
}

const (
	dateSuffixFormat = "2006-01-02-15-04-05"
	volumeCopyImage  = "alpine:3"
)

var copyLabelsToVolumeBackup = []string{
	"org.label-schema.version",
	"org.opencontainers.image.version",
}

var (
	// ErrContainerWait is returned on error response from container wait.
	ErrContainerWait = errors.New("ContainerWait")
	// ErrVolumeBackup is returned on error with volume backup.
	ErrVolumeBackup = errors.New("VolumeBackup")
	// ErrNotInstalledFromCli is returned when the current container was not installed via cli.
	ErrNotInstalledFromCli = errors.New("NotInstalledFromCli")
)

type upgradeResult struct{}

// Result is a command run result.
func (u *upgradeResult) Result() {}

// String stringifies command result.
func (u *upgradeResult) String() string {
	return "ok"
}

func NewUpgradeCommand(l *logrus.Entry, waitBeforeContainerStop time.Duration) *UpgradeCommand {
	return &UpgradeCommand{
		l:                       l,
		waitBeforeContainerStop: waitBeforeContainerStop,
	}
}

// RunCmdWithContext runs upgrade command.
func (c *UpgradeCommand) RunCmdWithContext(ctx context.Context, _ *flags.GlobalFlags) (commands.Result, error) { //nolint:unparam
	if c.l == nil {
		c.l = logrus.NewEntry(logrus.StandardLogger())
	}

	c.l.Info("Starting PMM Server upgrade via Docker")

	d, err := prepareDocker(ctx, c.docker, prepareOpts{install: false})
	if err != nil {
		return nil, err
	}
	c.docker = d

	currentContainer, err := c.docker.ContainerInspect(ctx, c.ContainerID)
	if err != nil {
		return nil, err
	}

	if !c.isInstalledViaCli(currentContainer) {
		if !c.confirmToContinue(c.ContainerID) {
			return nil, fmt.Errorf("%w: the existing PMM Server was not installed via pmm cli", ErrNotInstalledFromCli)
		}
	}

	c.l.Infof("Downloading PMM Server %s", c.DockerImage)
	if err := c.pullImage(ctx, c.DockerImage); err != nil {
		return nil, err
	}

	c.l.Infof("Stopping PMM Server in container %q", currentContainer.Name)

	// We wait a bit so the logs can be requested by pmm-managed before the container is stopped.
	select {
	case <-time.After(c.waitBeforeContainerStop):
	case <-ctx.Done():
		msg := "upgrade has been cancelled"
		if err = ctx.Err(); err != nil {
			msg = fmt.Sprintf(msg+". Error: %s", err)
		}

		return nil, errors.New(msg)
	}

	noTimeout := -1 * time.Second
	if err = c.docker.ContainerStop(ctx, currentContainer.ID, &noTimeout); err != nil {
		return nil, err
	}

	if err = c.backupVolumes(ctx, &currentContainer); err != nil {
		return nil, err
	}

	if err = c.runPMMServer(ctx, currentContainer); err != nil {
		return nil, err
	}

	// Disable restart policy in the old container
	_, err = c.docker.ContainerUpdate(ctx, currentContainer.ID, container.UpdateConfig{
		RestartPolicy: container.RestartPolicy{Name: "no"},
	})
	if err != nil {
		c.l.Info("We could not disable restart policy in the old container.")
		c.l.Infof(`We strongly recommend removing the old container manually with "docker rm %s"`, currentContainer.Name)
		c.l.Errorf("Error for reference: %#v", err)
	}

	c.l.Info("PMM Server update finished successfully")

	return &upgradeResult{}, nil
}

func (c *UpgradeCommand) isInstalledViaCli(container types.ContainerJSON) bool {
	if container.Config == nil {
		return false
	}

	v := container.Config.Labels["percona.pmm.source"]

	return v == "cli"
}

func (c *UpgradeCommand) confirmToContinue(containerID string) bool {
	actions := make([]string, 0, 4)
	actions = append(actions, fmt.Sprintf("- Stop the container %q", containerID))
	if !c.DisableBackup {
		actions = append(actions, fmt.Sprintf("- Back up all volumes in %q", containerID))
	}
	actions = append(actions, fmt.Sprintf("- Mount all volumes from %q in the new container", containerID))
	actions = append(actions, fmt.Sprintf("- Share the same network ports as in %q", containerID))

	fmt.Printf(`
PMM Server in the container %[1]q was not installed via pmm cli.
We will attempt to upgrade the container and perform the following actions:

%s

The container %[1]q will NOT be removed. You can remove it manually later, if needed.

`, containerID, strings.Join(actions, "\n"))

	if c.AssumeYes {
		return true
	}

	fmt.Print("Are you sure you want to continue? [y/N] ")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()

	return strings.ToLower(input) == "y"
}

func (c *UpgradeCommand) backupVolumes(ctx context.Context, container *types.ContainerJSON) error {
	if c.DisableBackup {
		c.l.Info("Backup of volumes is disabled. Skipping backup.")
		return nil
	}

	c.l.Info("Starting backup of volumes")

	c.l.Infof("Downloading %q", volumeCopyImage)
	if err := c.pullImage(ctx, volumeCopyImage); err != nil {
		return err
	}

	now := time.Now()
	for _, m := range container.Mounts {
		if m.Type != mount.TypeVolume {
			continue
		}

		// Copy labels from original container to backup volume
		labels := make(map[string]string, 1+len(copyLabelsToVolumeBackup))
		for _, l := range copyLabelsToVolumeBackup {
			if _, ok := container.Config.Labels[l]; ok {
				labels[l] = container.Config.Labels[l]
			}
		}

		labels["percona.pmm.created"] = now.Format(dateSuffixFormat)

		backupName := fmt.Sprintf("%s-backup-%s", m.Name, now.Format(dateSuffixFormat))
		_, err := c.docker.CreateVolume(ctx, backupName, labels)
		if err != nil {
			return err
		}

		c.l.Infof("Backing up volume %q to %q", m.Name, backupName)
		if err = c.backupVolumeViaContainer(ctx, m.Name, backupName); err != nil {
			return err
		}
	}

	return nil
}

func (c *UpgradeCommand) pullImage(ctx context.Context, imageName string) error {
	reader, err := c.docker.PullImage(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		c.l.Error(err)
	}

	return nil
}

func (c *UpgradeCommand) backupVolumeViaContainer(ctx context.Context, srcVolume, dstVolume string) error {
	c.l.Infof("Starting container to backup %q to %q", srcVolume, dstVolume)
	containerID, err := c.docker.RunContainer(ctx, &container.Config{
		Image: volumeCopyImage,
		Cmd:   strslice.StrSlice{"cp", "-prT", "/srv-original", "/srv-backup"},
		Labels: map[string]string{
			"percona.pmm.source": "cli-backup-container",
		},
	}, &container.HostConfig{
		Binds: []string{
			srcVolume + ":/srv-original:ro",
			dstVolume + ":/srv-backup:rw",
		},
		AutoRemove: true,
	}, "")
	if err != nil {
		return err
	}

	c.l.Info("Backing up volume data")
	waitC, errC := c.docker.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case res := <-waitC:
		if res.Error != nil {
			return fmt.Errorf("%w: %s", ErrContainerWait, res.Error.Message)
		}

		if res.StatusCode != 0 {
			return fmt.Errorf("%w: backup exited with code %d", ErrVolumeBackup, res.StatusCode)
		}
	case err := <-errC:
		return err
	}

	return nil
}

func (c *UpgradeCommand) runPMMServer(ctx context.Context, currentContainer types.ContainerJSON) error {
	c.l.Info("Starting PMM Server")

	containerName := fmt.Sprintf("%s-%s", c.NewContainerNamePrefix, time.Now().Format(dateSuffixFormat))
	if c.NewContainerName != "" {
		containerName = c.NewContainerName
	}

	containerID, err := startPMMServer(
		ctx, nil, currentContainer.ID, c.DockerImage,
		c.docker, currentContainer.HostConfig.PortBindings, containerName,
		currentContainer.Config.Env)
	if err != nil {
		return err
	}

	c.l.Debugf("Started PMM Server in container %q", containerID)

	return nil
}
