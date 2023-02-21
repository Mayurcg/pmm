// Copyright 2023 Percona LLC
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

// Package update implements update API.
package update

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm/admin/cli/flags"
	dockerCmd "github.com/percona/pmm/admin/commands/pmm/server/docker"
	"github.com/percona/pmm/admin/pkg/docker"
	"github.com/percona/pmm/api/updatepb"
)

// Make sure this does not include any directories, just a basename.
const logsFileNamePattern = "upgrade.*.log"

type Server struct {
	// Context coming from cli commands. When cancelled, the command has been cancelled.
	cliCtx             context.Context
	docker             containerManager
	dockerImage        string
	gRPCMessageMaxSize uint32

	updateInProgress map[string]struct{}
	updateMu         sync.RWMutex

	updatepb.UnimplementedUpdateServer
}

// New returns new instance of Server.
func New(ctx context.Context, dockerImage string, gRPCMessageMaxSize uint32) (*Server, error) {
	d, err := docker.New(nil)
	if err != nil {
		return nil, err
	}

	return &Server{
		cliCtx:             ctx,
		docker:             d,
		dockerImage:        dockerImage,
		gRPCMessageMaxSize: gRPCMessageMaxSize,
		updateInProgress:   make(map[string]struct{}, 1),
	}, nil
}

// StartUpdate starts PMM Server update.
func (s *Server) StartUpdate(ctx context.Context, req *updatepb.StartUpdateRequest) (*updatepb.StartUpdateResponse, error) {
	containerID := req.ContainerId

	logrus.Debugf("Inspecting container %s", containerID)
	container, err := s.docker.ContainerInspect(ctx, containerID)
	if err != nil {
		logrus.Errorf("Could not inspect container %s. Error: %v", containerID, err)
		return nil, err
	}

	if !container.State.Running {
		return nil, fmt.Errorf("container %s is not running", containerID)
	}

	logFile, err := os.CreateTemp("", logsFileNamePattern)
	if err != nil {
		return nil, err
	}

	go func() {
		logrus.Infof("Starting update of container %s with log file %s", containerID, logFile.Name())
		defer func() {
			// Keep the log around for a short period of time
			<-time.After(5 * time.Minute)
			if err := os.Remove(logFile.Name()); err != nil {
				logrus.Error(err)
			}
		}()

		logger := logrus.New()
		logger.SetOutput(io.MultiWriter(logFile, os.Stdout))

		cmd := dockerCmd.NewUpgradeCommand(
			logger.WithField("update", logFile.Name()),
			5*time.Second)
		cmd.ContainerID = containerID
		cmd.DockerImage = s.dockerImage
		cmd.NewContainerNamePrefix = "pmm-server"

		// Store update in progress info.
		s.updateMu.Lock()
		s.updateInProgress[logFile.Name()] = struct{}{}
		s.updateMu.Unlock()

		defer func() {
			s.updateMu.Lock()
			defer s.updateMu.Unlock()

			delete(s.updateInProgress, logFile.Name())
		}()

		_, err := cmd.RunCmdWithContext(s.cliCtx, &flags.GlobalFlags{})
		if err != nil {
			logger.Errorf("Could not update container %s. Error: %v", containerID, err)
		}
	}()

	return &updatepb.StartUpdateResponse{LogsToken: logFile.Name()}, nil
}

// UpdateStatus returns PMM Server update status.
func (s *Server) UpdateStatus(ctx context.Context, req *updatepb.UpdateStatusRequest) (*updatepb.UpdateStatusResponse, error) { //nolint:unparam
	var err error
	var lines []string
	var newOffset uint32
	var done bool

	// wait up to 30 seconds for new log lines
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for ctx.Err() == nil {
		done = !s.isUpdateRunning(req.LogsToken)
		if done {
			// give a second to flush logs to file
			time.Sleep(time.Second)
		}

		lines, newOffset, err = s.getLogs(req.LogsToken, req.Offset)
		if err != nil {
			logrus.Warn(err)
		}

		if len(lines) != 0 || done {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	if err = ctx.Err(); !errors.Is(err, context.DeadlineExceeded) {
		logrus.Warnf("context error during UpdateStatus: %s", err)
	}

	return &updatepb.UpdateStatusResponse{
		Lines:  lines,
		Offset: newOffset,
		Done:   done,
	}, nil
}

func (s *Server) isUpdateRunning(name string) bool {
	s.updateMu.RLock()
	defer s.updateMu.RUnlock()

	_, ok := s.updateInProgress[name]

	return ok
}

// IsAnyUpdateRunning returns true if there's at least one update in progress.
func (s *Server) IsAnyUpdateRunning() bool {
	s.updateMu.RLock()
	defer s.updateMu.RUnlock()

	return len(s.updateInProgress) != 0
}

// getLogs returns some lines and a new offset from a log file starting from the given offset.
// It may return zero lines and the same offset. Caller is expected to handle this.
func (s *Server) getLogs(filePath string, offset uint32) ([]string, uint32, error) {
	if err := s.isValidLogPath(filePath); err != nil {
		return nil, 0, errors.WithStack(err)
	}

	f, err := os.Open(filePath) //nolint:gosec
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	defer f.Close() //nolint:gosec

	if _, err = f.Seek(int64(offset), io.SeekStart); err != nil {
		return nil, 0, errors.WithStack(err)
	}

	lines := make([]string, 0, 10)
	reader := bufio.NewReader(f)
	newOffset := offset
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			newOffset += uint32(len(line))
			if newOffset-offset > s.gRPCMessageMaxSize {
				return lines, newOffset - uint32(len(line)), nil
			}
			lines = append(lines, strings.TrimSuffix(line, "\n"))
			continue
		}
		if err == io.EOF {
			err = nil
		}
		return lines, newOffset, errors.WithStack(err)
	}
}

func (s *Server) isValidLogPath(filePath string) error {
	filename := path.Base(filePath)
	match, err := path.Match(logsFileNamePattern, filename)
	if err != nil {
		return err
	}

	if !match {
		return fmt.Errorf("invalid log file path provided")
	}

	return nil
}
