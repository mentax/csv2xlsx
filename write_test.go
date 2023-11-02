package main

import (
	"github.com/tealeg/xlsx/v3"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
}
