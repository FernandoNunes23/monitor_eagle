package client

import (
  "github.com/spf13/cobra"
  "client/config"
)

func showMemory() {
	conf     := config.InitConfig()
	scheme 	 := "http://"
	hostname := conf.Hostname
	uri		 := ":8000/memory"

	url := scheme + hostname + uri
    PrintResponseBody(url)
}

var show = &cobra.Command{
  Use:   "show",
  Short: "Show memory info",
  Long:  `All software has versions. This is Eagle's`,
  Run: func(cmd *cobra.Command, args []string) {
      showMemory();
  },
}