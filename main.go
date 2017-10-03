package main

import (
	"fmt"
	"os"

	"encoding/csv"
	"io"
	"strconv"

	"github.com/tealeg/xlsx"
	"github.com/urfave/cli"
)

const SHEET_NAME_TEMPLATE = "Sheet %i"

func main() {

	initCommandLine(os.Args)
}

func writeAllSheets(xlFile *xlsx.File, dataFiles []string, sheetNames []string, exampleRowNumber int) (err error) {

	for i, dataFileName := range dataFiles {

		sheetName := sheetNames[i]
		if sheetName == "" {
			sheetName = fmt.Sprintf(SHEET_NAME_TEMPLATE, i)
		}

		sheet, ok := xlFile.Sheet[sheetName]
		if ok != true {
			sheet, err = xlFile.AddSheet(sheetName)

			if err != nil {
				return err
			}
		}

		var exampleRow *xlsx.Row
		if exampleRowNumber != 0 {
			exampleRow = sheet.Rows[exampleRowNumber]
		}

		err = writeSheet(dataFileName, sheet, exampleRow)

		if exampleRow != nil {
			// remove example row
			sheet.Rows = append(sheet.Rows[:exampleRowNumber], sheet.Rows[exampleRowNumber+1:]...)
		}
	}

	return err
}

func writeSheet(dataFileName string, sheet *xlsx.Sheet, exampleRow *xlsx.Row) error {

	data := getCsvData(dataFileName)

	var i int
	for {
		record, err := data.Read()

		if err == io.EOF || record == nil {
			break
		} else if err != nil {
			return err
		}

		//if i > 5000 {
		//	break
		//}

		//if i%500 == 0 {
		//	fmt.Println(i)
		//}

		i++

		writeRowToXls(sheet, record, exampleRow)
	}

	return nil
}

func buildXls(c *cli.Context, p *params) (err error) {

	var xlFile *xlsx.File
	if p.xlsxTemplate == "" {
		xlFile = xlsx.NewFile()
	} else {
		xlFile, err = xlsx.OpenFile(p.xlsxTemplate)
		if err != nil {
			return err
		}
	}

	writeAllSheets(xlFile, p.input, p.sheets, p.row)

	return xlFile.Save(p.output)
}

func writeRowToXls(sheet *xlsx.Sheet, record []string, exampleRow *xlsx.Row) {

	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	//row.WriteSlice( &record , -1)

	for k, v := range record {
		cell = row.AddCell()

		setCellValue(cell, v)

		if exampleRow != nil {
			style := exampleRow.Cells[k].GetStyle()

			cell.SetStyle(style)
		}
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
