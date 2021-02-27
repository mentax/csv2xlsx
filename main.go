package main

import (
	"encoding/csv"
	"errors"
	"os"
)

// SheetNamesTemplate define name's for new created sheets.
var SheetNamesTemplate = "Sheet %d"

func main() {
	initCommandLine(os.Args)
}

// getCsvData read's data from CSV file.
func getCsvData(dataFileName string) (*csv.Reader, error) {
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		return nil, errors.New("Problem with reading data from " + dataFileName)
	}

	return csv.NewReader(dataFile), nil
}
