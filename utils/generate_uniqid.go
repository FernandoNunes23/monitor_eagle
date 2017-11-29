package utils

import (
	"github.com/sony/sonyflake"
	"log"
)

func GenerateUniqid() (uint64, error) {
	return genSonyflake()
}

func genSonyflake() (uint64, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
		return 0, err 
	}
	// Note: this is base16, could shorten by encoding as base62 string
	return id, err
} 