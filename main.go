package main

import (
	"fmt"
	"os"

	"encoding/csv"
	"io"
	"strconv"

	"github.com/tealeg/xlsx"
)

var exampleRowNumber = 3

func init() {

}

func main() {

	initCommandLine(os.Args)
	return

	dataFileName := "example/test.csv"

	r := getCsvData(dataFileName)

	//fmt.Println(dataFileName)

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

		setCellValue(cell, v)

		style := exampleRow.Cells[k].GetStyle()

		cell.SetStyle(style)
	}
}

// setCellValue set data in correct format.
func setCellValue(cell *xlsx.Cell, v string) {
	intVal, err := strconv.Atoi(v)
	if err == nil {
		if intVal < 100000000000 { // Long numbers are displayed incorrectly in Excel
			cell.SetInt(intVal)
			return
		}
		cell.Value = v
		return
	}

	floatVal, err := strconv.ParseFloat(v, 64)
	if err == nil {
		cell.SetFloat(floatVal)
		return
	}
	cell.Value = v
}

// getCsvData read's data from CSV file.
func getCsvData(dataFileName string) *csv.Reader {
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csv.NewReader(dataFile)
}
