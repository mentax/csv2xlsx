package main

import (
	"fmt"
	"os"

	"encoding/csv"
	"github.com/tealeg/xlsx"
	"io"
	"strconv"
)

var exampleRowNumber = 3

func main() {
	initCommandLine(os.Args)

	return

	dataFileName := "example/test.csv"

	r := getSheetData(dataFileName)

	fmt.Println(dataFileName)

	xlFile, err := xlsx.OpenFile("./example/default.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sheet := xlFile.Sheet["BO"]
	exampleRow := sheet.Rows[exampleRowNumber]

	var i int
	for {
		record, err := r.Read()

		if err == io.EOF || record == nil {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		//if i > 5000 {
		//	break
		//}

		if i%500 == 0 {
			fmt.Println(i)
		}

		i++

		writeRowToXls(sheet, record, exampleRow)
	}

	// remove example row
	sheet.Rows = append(sheet.Rows[:exampleRowNumber], sheet.Rows[exampleRowNumber+1:]...)

	err = xlFile.Save("./example/Result.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeRowToXls(sheet *xlsx.Sheet, record []string, exampleRow *xlsx.Row) {

	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	//row.WriteSlice( &record , -1)

	for k, v := range record {
		cell = row.AddCell()

		intVal, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			cell.SetInt64(intVal)
		} else {
			floatVal, err := strconv.ParseFloat(v, 64)

			if err == nil {
				cell.SetFloat(floatVal)
			} else {
				cell.Value = v
			}
		}

		style := exampleRow.Cells[k].GetStyle()

		cell.SetStyle(style)
	}
}

func getSheetData(dataFileName string) *csv.Reader {
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csv.NewReader(dataFile)
}
