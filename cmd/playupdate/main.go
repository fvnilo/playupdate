package main

import (
	"log"
	"os"
	"runtime"

	"github.com/nylo-andry/playupdate/csv"
	"github.com/nylo-andry/playupdate/http"
	"github.com/nylo-andry/playupdate/worker"
)

func main() {
	f, err := os.Open("examples/input.csv")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	macAddresses, err := csv.ReadFile(f)
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}

	service := http.NewPlayUpdateService("http://localhost:8080")
	pool := worker.NewPool(runtime.NumCPU(), service)
	pool.Start(macAddresses)
}
