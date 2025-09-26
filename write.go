package main

import (
	"fmt"
	"io"
	"unicode"

	"codeberg.org/tealeg/xlsx/v4"
)

func writeAllSheets(xlFile *xlsx.File, dataFiles []string, sheetNames []string, exampleRowNumber int, delimiter rune, startFrom int) (err error) {

	for i, dataFileName := range dataFiles {

		sheet, err := getSheet(xlFile, sheetNames, i)
		if err != nil {
			return err
		}

		var exampleRow *xlsx.Row
		if exampleRowNumber != 0 && exampleRowNumber <= sheet.MaxRow {
			// example row counting from 1
			exampleRow, _ = sheet.Row(exampleRowNumber - 1)

			err = sheet.RemoveRowAtIndex(exampleRowNumber - 1)
			if err != nil {
				return err
			}
		}

		err = writeSheet(dataFileName, sheet, exampleRow, delimiter, startFrom)

		if err != nil {
			return err
		}
	}

	return nil
}

func writeSheet(dataFileName string, sheet *xlsx.Sheet, exampleRow *xlsx.Row, delimiter rune, startFrom int) error {

	data, err := getCsvData(dataFileName)

	if err != nil {
		return err
	}

	data.Comma = delimiter

	var i int
	for {
		record, err := data.Read()

		if err == io.EOF || record == nil {
			break
		} else if err != nil {
			return err
		}

		i++

		if i <= startFrom {
			continue
		}

		// if i > 5000 {
		//	break
		// }

		// if i%500 == 0 {
		// 	fmt.Println(i)
		// }

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

	err = writeAllSheets(xlFile, p.input, p.sheets, p.exampleRow, p.delimiter, p.startFrom)
	if err != nil {
		return err
	}

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
			format := exampleRow.GetCell(k).GetNumberFormat()

			cell.SetStyle(style)
			cell.SetFormat(format)
		}
	}
}

// setCellValue set data in correct format.
func setCellValue(cell *xlsx.Cell, v string) {

	// fast path
	if v == "" {
		cell.SetString("")
		return
	}

	// field marked as string
	if v[0] == '\'' {
		cell.SetString(v[1:])
		return
	}

	if isNumeric(v) {
		cell.SetNumeric(v)
		return
	}

	cell.SetString(v)
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

func isNumeric(s string) bool {
	if s == "" {
		return false
	}

	var haveDot, haveDigit bool
	for i, c := range s {
		switch {
		case c == '.':
			if haveDot {
				return false
			}
			haveDot = true
		case c == '-':
			if i > 0 {
				return false
			}
		case unicode.IsDigit(c):
			haveDigit = true
		default:
			return false
		}
	}

	return haveDigit
}
