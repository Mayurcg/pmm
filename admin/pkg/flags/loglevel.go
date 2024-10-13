// Copyright (C) 2023 Percona LLC
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

// Package flags provides common flags and helpers for it.
package flags

import (
	"fmt"
	"strings"

	"github.com/AlekSi/pointer"
)

// LogLevelFatalFlags contains log level flag with "fatal" option.
type LogLevelFatalFlags struct {
	LogLevel LogLevel `name:"log-level" enum:"debug,info,warn,error,fatal" default:"warn" help:"Service logging level. One of: [${enum}]. Default: ${default}"`
}

// LogLevelNoFatalFlags contains log level flag without "fatal" option.
type LogLevelNoFatalFlags struct {
	LogLevel LogLevel `name:"log-level" enum:"debug,info,warn,error" default:"warn" help:"Service logging level. One of: [${enum}]. Default: ${default}"`
}

type LogLevel string

// EnumValue returns pointer to string representation of LogLevel.
func (l LogLevel) EnumValue() *string {
	return pointer.To(convertEnum("LOG_LEVEL", string(l)))
}

func convertEnum(prefix string, value string) string {
	return fmt.Sprintf("%s_%s", prefix, strings.ToUpper(value))
}
