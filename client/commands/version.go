package client

import (
	"github.com/spf13/cobra"
	"fmt"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Monitor Eagle",
  Long:  `All software has versions. This is Eagle's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Eagle Version 0.1 -- HEAD")
  },
}