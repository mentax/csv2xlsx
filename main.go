package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"os"
)

// SheetNamesTemplate define name's for new created sheets.
var SheetNamesTemplate = "Sheet %d"

func main() {
	initCommandLine(os.Args)
}

func getSheet(xlFile *xlsx.File, sheetNames []string, i int) (sheet *xlsx.Sheet, err error) {
	var sheetName string
	if len(sheetNames) > i {
		sheetName = sheetNames[i]
	} else {
		sheetName = fmt.Sprintf(SheetNamesTemplate, i+1)
	}

	sheet, ok := xlFile.Sheet[sheetName]
	if ok != true {
		sheet, err = xlFile.AddSheet(sheetName)

		if err != nil {
			return nil, err
		}
	}
	return sheet, nil
}

// getCsvData read's data from CSV file.
func getCsvData(dataFileName string) (*csv.Reader, error) {
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		return nil, errors.New("Problem with reading data from " + dataFileName)
	}

	return csv.NewReader(dataFile), nil
}
