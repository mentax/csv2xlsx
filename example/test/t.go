package main

import (
	"fmt"

	"codeberg.org/tealeg/xlsx/v4"
)

func main() {

	var xlsxFile = "./excel.xlsx"

	xlFile, err := xlsx.OpenFile(xlsxFile)
	if err != nil {
		panic(err)
	}

	var sheetName = "Sheet 1"
	sheet, ok := xlFile.Sheet[sheetName]

	if !ok {
		panic(" no sheet")
	}

	cell, _ := sheet.Cell(0, 1)

	fmt.Printf("\n %+v ", *cell)

	cell, _ = sheet.Cell(1, 1)

	fmt.Printf("\n %+v ", *cell)

	cell, _ = sheet.Cell(2, 1)

	fmt.Printf("\n %+v ", *cell)
}
