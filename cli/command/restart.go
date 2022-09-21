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
	"github.com/opencurve/pigeon/internal/core"
	cliutils "github.com/opencurve/pigeon/internal/utils"
	"github.com/spf13/cobra"
)

type restartOptions struct {
}

func NewRestartCommand(pigeon *core.Pigeon) *cobra.Command {
	var options restartOptions

	cmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart pigeon",
		Args:  cliutils.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRestart(pigeon, options)
		},
		DisableFlagsInUseLine: true,
	}
	return cmd
}

func runRestart(pigeon *core.Pigeon, options restartOptions) error {
	return nil
}
