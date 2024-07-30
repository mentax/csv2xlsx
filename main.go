package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

// SheetNamesTemplate define name's for new created sheets.
var SheetNamesTemplate = "Sheet %d"

func main() {
	if err := initCommandLine(os.Args); err != nil {
		log.Fatal(err)
	}
}

// getCsvData read's data from CSV file.
func getCsvData(dataFileName string) (*csv.Reader, error) {
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		return nil, errors.New("Can not read data from " + dataFileName)
	}

	r := csv.NewReader(dataFile)
	r.ReuseRecord = true
	return r, nil
}
