package csv

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		macAddresses = append(macAddresses, record[0])
	}

	return macAddresses, nil
}
