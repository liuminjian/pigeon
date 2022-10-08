/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: Pigeon
 * Created Date: 2022-09-21
 * Author: Jingli Chen (Wine93)
 */

package command

import (
	"fmt"
	"syscall"

	"github.com/opencurve/pigeon/internal/core"
	cliutils "github.com/opencurve/pigeon/internal/utils"
	"github.com/spf13/cobra"
)

type stopOptions struct {
	filename string
}

func NewStopCommand(pigeon *core.Pigeon) *cobra.Command {
	var options stopOptions

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop pigeon",
		Args:  cliutils.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStop(pigeon, options)
		},
		DisableFlagsInUseLine: true,
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.filename, "conf", "c", pigeon.DefaultConfFile(), "Specify pigeon configure file")

	return cmd
}

func runStop(pigeon *core.Pigeon, options stopOptions) error {
	pid, err := getPid(pigeon, options.filename)
	if err != nil {
		return fmt.Errorf("read pid file failed: %s", err)
	}
	return syscall.Kill(pid, syscall.SIGTERM)
}
