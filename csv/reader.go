package csv

import (
	"encoding/csv"
	"io"
)

func ReadFile(r io.Reader) ([]string, error) {
	csvReader := csv.NewReader(r)

	macAddresses := make([]string, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		macAddresses = append(macAddresses, record[0])
	}

	if len(macAddresses) == 0 {
		return macAddresses, nil
	}

	return macAddresses[1:], nil
}
