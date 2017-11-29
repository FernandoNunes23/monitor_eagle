package parser

import (
	"strconv"
	"log"
)

func GetUsedAmount(total string, free string) int {
	var err error

	total_int, err := strconv.Atoi(total)
	if err != nil {
        log.Fatal(err)
    }

	free_int, err  := strconv.Atoi(free)
	if err != nil {
        log.Fatal(err)
    }


	return (total_int - free_int)
}