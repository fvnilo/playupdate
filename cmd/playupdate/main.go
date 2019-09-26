package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("examples/input.csv")
	if err != nil {
		log.Fatalf("Could not open input file: %s", err)
	}
	defer f.Close()
	r := csv.NewReader(f)

	macAddresses := make([]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Parsing file failed: %s", err)
		}

		macAddresses = append(macAddresses, record[0])
	}

	var wg sync.WaitGroup
	inputs := make(chan string)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(inputs, &wg, "http://localhost:8080")
	}

	for _, mac := range macAddresses[1:] {
		inputs <- mac
	}
	close(inputs)

	wg.Wait()
}

func worker(inputs <-chan string, wg *sync.WaitGroup, baseUrl string) {
	defer wg.Done()

	client := &http.Client{}
	for mac := range inputs {
		url := fmt.Sprintf("%s/%s:%s", baseUrl, "/profiles/clienId", mac)
		req, _ := http.NewRequest(http.MethodPut, url, strings.NewReader("any thing"))
		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			defer res.Body.Close()
			fmt.Printf("Done processing mac %s\n", mac)
		}
	}
}
