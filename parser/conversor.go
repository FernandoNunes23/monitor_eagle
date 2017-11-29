package parser

import (
	"strconv"
	"log"
)

func ConvertKbToMb(kbValue string) int {
	num, err := strconv.Atoi(kbValue)

	if err != nil {
        log.Fatal(err)
    }

	mbValue := num / 1024

	return mbValue
}