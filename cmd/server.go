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
