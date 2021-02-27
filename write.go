package main

import (
	"github.com/tealeg/xlsx/v3"
	"io"
	"math"
	"strconv"
)

func writeAllSheets(xlFile *xlsx.File, dataFiles []string, sheetNames []string, exampleRowNumber int) (err error) {

	for i, dataFileName := range dataFiles {

		sheet, err := getSheet(xlFile, sheetNames, i)
		if err != nil {
			return err
		}

		var exampleRow *xlsx.Row
		if exampleRowNumber != 0 && exampleRowNumber <= sheet.MaxRow {
			// example row counting from 1
			exampleRow, _ = sheet.Row(exampleRowNumber - 1)

			sheet.RemoveRowAtIndex(exampleRowNumber - 1)
		}

		err = writeSheet(dataFileName, sheet, exampleRow)

		if err != nil {
			return err
		}
	}

	return nil
}

func writeSheet(dataFileName string, sheet *xlsx.Sheet, exampleRow *xlsx.Row) error {

	data, err := getCsvData(dataFileName)

	if err != nil {
		return err
	}

	var i int
	for {
		record, err := data.Read()

		if err == io.EOF || record == nil {
			break
		} else if err != nil {
			return err
		}

		// if i > 5000 {
		//	break
		// }

		// if i%500 == 0 {
		// 	fmt.Println(i)
		// }

		i++

		writeRowToXls(sheet, record, exampleRow)
	}

	return nil
}

func buildXls(p *params) (err error) {

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

	var cellsLen int
	if exampleRow != nil {
		cellsLen = exampleRow.Sheet.MaxCol
	}

	for k, v := range record {
		cell = row.AddCell()

		setCellValue(cell, v)

		if exampleRow != nil && cellsLen > k {
			style := exampleRow.GetCell(k).GetStyle()

			cell.SetStyle(style)
		}
	}
}

// setCellValue set data in correct format.
func setCellValue(cell *xlsx.Cell, v string) {
	intVal, err := strconv.Atoi(v)
	if err == nil {
		if intVal < math.MinInt32 { // Long numbers are displayed incorrectly in Excel
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
