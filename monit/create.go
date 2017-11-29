package monit

import (
	"utils"
	"types"
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var pathToMount string
var	configFile	string

func Init(path string) {
	pathToMount = path
}

func Create(m types.Monit) bool {
	
	if mountStructure(pathToMount, m) == false {
		return false
	}

	return true
}	

func mountStructure(path string, config types.Monit) bool{
	if utils.VerifyIfExist(path) == false {
		os.Mkdir(path, 0777)
	}

	if utils.VerifyIfExist(path + "/" + config.Name) == false {
		os.Mkdir(path + "/" + config.Name, 0777)
	}

	if utils.VerifyIfExist(path + "/" + config.Name + "/config.json") == false {
		os.Create(path + "/" + config.Name + "/config.json")

		configFile = path + "/" + config.Name + "/config.json"

		c, err := json.Marshal(config); if err != nil {
			fmt.Print(err)
		} 

		ioutil.WriteFile(configFile, c, 0644)
	}

	return true
}