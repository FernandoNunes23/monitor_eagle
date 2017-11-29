package client

import (
	"github.com/spf13/cobra"
  "api/server"
)

var RootCmd = &cobra.Command{
  Use:   "eagle",
  Short: "Eagle is a simple way to create your monit",
  Long: `Eagle`,
  Run: func(cmd *cobra.Command, args []string) {
    server.Init()
  },
}

func init() {
  RootCmd.AddCommand(versionCmd)
  RootCmd.AddCommand(show)
}