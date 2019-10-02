package main

import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/nylo-andry/playupdate/csv"
	"github.com/nylo-andry/playupdate/http"
	"github.com/nylo-andry/playupdate/worker"

	"github.com/urfave/cli"
)

func main() {
	var input, baseUrl string
	app := cli.NewApp()
	app.Name = "playupdate"
	app.Usage = "Update music players in batch"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Load input from `FILE` (mandatory)",
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "api-url",
			Usage:       "Api base `URL` (optional, defaults to http://localhost:8080)",
			Destination: &baseUrl,
		},
	}
	app.Action = func(c *cli.Context) error {
		if input == "" {
			return errors.New("No input file specified")
		}
		f, err := os.Open(input)
		if err != nil {
			log.Fatalf("Could not open file: %s", err)
		}
		defer f.Close()

		macAddresses, err := csv.ReadFile(f)
		if err != nil {
			log.Fatalf("Could not open file: %s", err)
		}

		if baseUrl == "" {
			baseUrl = "http://localhost:8080"
		}
		service := http.NewPlayUpdateService(baseUrl)
		pool := worker.NewPool(runtime.NumCPU(), service)
		pool.Start(macAddresses)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
