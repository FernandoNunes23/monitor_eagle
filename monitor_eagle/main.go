package main

/*import (
	"os"
	"log"
	"parser"
	"logger"
	"formatter"
	"api/server"
	"commands"
)
func main() {
	logger.Init()

	file, err := os.Open("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    mem := parser.Parse(file)
   
    data := formatter.FormatMemoryLog(mem)

    server.Test()

	logger.Log(data)
}*/

import (
  "fmt"
  "os"
  "client/commands"
)

func main() {
  if err := client.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}