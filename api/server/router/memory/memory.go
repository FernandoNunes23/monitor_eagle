package memory

import (
	"parser"
	"os"
	"log"
	"net/http"
	"encoding/json"
)

func GetMemoryInfo(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    mem := parser.ParseMemory(file)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(mem)
}

func GetSwapInfo(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    swap := parser.ParseSwap(file)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(swap)
}