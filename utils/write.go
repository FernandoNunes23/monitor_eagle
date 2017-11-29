package utils

import (
	"os"
	"fmt"
)

func WriteFile(data string, path string) {
	
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	
	if isError(err) { return }
	
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	if isError(err) { return }

	// save changes
	err = file.Sync()
	if isError(err) { return }
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}