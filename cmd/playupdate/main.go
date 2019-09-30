package main

import (
	"log"
	"runtime"

	"github.com/nylo-andry/playupdate/csv"
	"github.com/nylo-andry/playupdate/http"
	"github.com/nylo-andry/playupdate/worker"
)

func main() {
	macAddresses, err := csv.ReadFile("examples/input.csv")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}

	service := http.NewPlayUpdateService("http://localhost:8080")
	pool := worker.NewPool(runtime.NumCPU(), service)
	pool.Start(macAddresses)
}
