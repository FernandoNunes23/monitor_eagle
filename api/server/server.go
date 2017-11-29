package server

import (
	"net/http"
	"log"
	"api/server/router"
	"github.com/sevlyar/go-daemon"
)

func Init() {
	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	r := router.NewRouter()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}