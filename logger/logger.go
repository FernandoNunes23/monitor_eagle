package logger

import (
	"os"
	"log"
	"bytes"
	"time"
	"utils"
)

var buf    bytes.Buffer
var _logger = log.New(&buf, "INFO: ", log.Lshortfile)
var filename string
var datetime string
var pathname string 

func Init(path string){
	t := time.Now()
	datetime = t.Format("2006-01-02 15:04:05")
	filename = t.Format("2006-01-02") + ".log"
	pathname = path
}

func Log(data string) {

	if utils.VerifyIfExist(pathname + "/logs") == false {
		os.Mkdir(pathname + "/logs", 0777)
	}

	path := pathname + "/logs/" + filename

	if utils.VerifyIfExist(path) == false {
		os.Create(path)
	}

	logLine := datetime + " || " + data

	utils.WriteFile(logLine, path)
}