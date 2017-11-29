package utils

import (
	"os"
)

func VerifyIfExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
	  return false
	}

	return true
}