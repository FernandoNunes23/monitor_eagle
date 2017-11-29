package monit

import (
	"monit"
	"net/http"
	"types"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"logger"
	"os"
	"log"
	"parser"
	"time"
)

func CreateMonit(w http.ResponseWriter, r *http.Request) {
	config := types.MonitConfig{}

	config.Type 	= r.PostFormValue("configType")
	config.Resource = r.PostFormValue("configResource")
	config.Time     = r.PostFormValue("configTime")

	createMonit(r.PostFormValue("name"), config)

	w.Write([]byte("Monit Created!"))
}

func GetMonitInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
  	m := getMonitInfo(vars["name"])

	json.NewEncoder(w).Encode(m)
}

func ShowCreatedMonits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	var Monits []types.Monit
	dirs, err := ioutil.ReadDir("/tmp/eagle/")

	if err != nil {
		fmt.Print(err)
	} 
   	
   	for _, dir := range dirs {
   		mon := getMonitInfo(dir.Name())
   		Monits = append(Monits, mon)
   	}

   	json.NewEncoder(w).Encode(Monits)
}

func LaunchMonit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	launchMonit(vars["name"])
}

func launchMonit(name string) {
	m := getMonitInfo(name)

	if m.Config.Resource == "memory" {
		go startLoggingMemory(m)
	}
}

func startLoggingMemory(m types.Monit) {
	duration := time.Duration(5 * time.Second)
	ticker := time.NewTicker(duration)
	for t := range ticker.C {
		mem := getMemoryInfo()
		logger.Init("/tmp/eagle/" + m.Name)
		
		memJson, err := json.Marshal(mem); if err != nil {
			fmt.Print(err)
		}

		s := string(memJson[:])
		fmt.Println("Logging at", t)
		logger.Log(s)
    }
}

func createMonit(name string, config types.MonitConfig) {
	monit.NewMonit(name, config)
}

func getMonitInfo(name string) types.Monit {
	var m types.Monit 
	m = types.Monit{}

	filePath := "/tmp/eagle/" + name +"/config.json" 

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Print(err)
	}
   	
   	json.Unmarshal(file, &m)

	return m
}

func getMemoryInfo() types.Memory {
	file, err := os.Open("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    mem := parser.ParseMemory(file)

	return mem
}