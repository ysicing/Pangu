// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"gitea.ysicing.net/cloud/pangu/pkg/server"
	"github.com/spf13/cobra"
)

func serverCommand() *cobra.Command {
	s := &cobra.Command{
		Use:   "server",
		Short: "core server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Serve()
		},
	}
	return s
}
