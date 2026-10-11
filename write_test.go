package main

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"codeberg.org/tealeg/xlsx/v4"

	"github.com/stretchr/testify/assert"
)

func TestWriteSheetCSVErrors(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter rune
		startFrom int
		rows      int
		cause     error
	}{
		{name: "first missing quote", input: "\"unterminated\n", delimiter: ',', cause: csv.ErrQuote},
		{name: "first bare quote", input: "bad\"value\n", delimiter: ',', cause: csv.ErrBareQuote},
		{name: "later missing quote", input: "name\n\"unterminated\n", delimiter: ',', rows: 1, cause: csv.ErrQuote},
		{name: "partial record", input: "name,value\nvalid,\"unterminated\n", delimiter: ',', rows: 1, cause: csv.ErrQuote},
		{name: "unequal fields", input: "name,value\none\n", delimiter: ',', rows: 1, cause: csv.ErrFieldCount},
		{name: "skipped malformed row", input: "\"unterminated\n", delimiter: ',', startFrom: 1, cause: csv.ErrQuote},
		{name: "zero delimiter", input: "name\n", delimiter: 0},
		{name: "quote delimiter", input: "name\n", delimiter: '"'},
		{name: "newline delimiter", input: "name\n", delimiter: '\n'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.csv")
			if err := os.WriteFile(path, []byte(tt.input), 0600); err != nil {
				t.Fatal(err)
			}
			file := xlsx.NewFile()
			sheet, err := file.AddSheet("Sheet")
			if err != nil {
				t.Fatal(err)
			}
			err = writeSheet(path, sheet, nil, tt.delimiter, tt.startFrom)
			assert.Error(t, err)
			if tt.cause != nil {
				assert.True(t, errors.Is(err, tt.cause), "got %v, want %v", err, tt.cause)
			}
			assert.Equal(t, tt.rows, sheet.MaxRow)
		})
	}
}

func TestWriteSheetCSVSuccess(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter rune
		startFrom int
		want      [][]string
	}{
		{name: "empty", delimiter: ','},
		{name: "blank lines", input: "\n\n", delimiter: ','},
		{name: "quoted fields", input: "name,value\n\"a,b\",\"line\nvalue\"\n", delimiter: ',', want: [][]string{{"name", "value"}, {"a,b", "line\nvalue"}}},
		{name: "semicolon", input: "name;value\n", delimiter: ';', want: [][]string{{"name", "value"}}},
		{name: "skipped header", input: "name\nvalue\n", delimiter: ',', startFrom: 1, want: [][]string{{"value"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.csv")
			if err := os.WriteFile(path, []byte(tt.input), 0600); err != nil {
				t.Fatal(err)
			}
			file := xlsx.NewFile()
			sheet, err := file.AddSheet("Sheet")
			if err != nil {
				t.Fatal(err)
			}
			assert.NoError(t, writeSheet(path, sheet, nil, tt.delimiter, tt.startFrom))
			assert.Equal(t, len(tt.want), sheet.MaxRow)
			for i, fields := range tt.want {
				row, err := sheet.Row(i)
				if err != nil {
					t.Fatal(err)
				}
				for j, field := range fields {
					assert.Equal(t, field, row.GetCell(j).Value)
				}
			}
		})
	}
}

func TestStrangeBehaviorOnInfValue(t *testing.T) {
	cell := &xlsx.Cell{}

	setCellValue(cell, "Inf")
	assert.Equal(t, "Inf", cell.Value)
	assert.Equal(t, xlsx.CellTypeString, cell.Type())
}

func TestValuesTypes(t *testing.T) {
	cell := &xlsx.Cell{}

	setCellValue(cell, "123")
	assert.Equal(t, "123", cell.Value)
	assert.Equal(t, xlsx.CellTypeNumeric, cell.Type())

	setCellValue(cell, "0.123")
	assert.Equal(t, "0.123", cell.Value)
	assert.Equal(t, xlsx.CellTypeNumeric, cell.Type())

	setCellValue(cell, "abc")
	assert.Equal(t, "abc", cell.Value)
	assert.Equal(t, xlsx.CellTypeString, cell.Type())
}

func TestInNumeric(t *testing.T) {

	assert.True(t, isNumeric("1.23"))
	assert.True(t, isNumeric("123"))
	assert.True(t, isNumeric("0.123"))
	assert.True(t, isNumeric("-0.123"))
	assert.True(t, isNumeric("-123"))

	assert.False(t, isNumeric("33-3"))
	assert.False(t, isNumeric("abc"))
	assert.False(t, isNumeric("Inf"))
	assert.False(t, isNumeric("1.."))
	assert.False(t, isNumeric("."))
	assert.False(t, isNumeric("-"))
}
