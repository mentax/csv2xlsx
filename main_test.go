package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const fileName = "./testFile.xlsx"

var p = &params{
	input:        []string{"./example/data.csv"},
	xlsxTemplate: "./example/template.xlsx",
	output:       fileName,
}

func TestBuildXls(t *testing.T) {
	os.Remove(fileName)

	err := buildXls(p)
	assert.Nil(t, err)

	assert.FileExists(t, fileName)

	f, err := os.Stat(fileName)
	assert.Nil(t, err)

	assert.Less(t, int64(8000), f.Size())

	os.Remove(fileName)
}

func BenchmarkBuildXls(b *testing.B) {
	os.Remove(fileName)

	for n := 0; n < b.N; n++ {
		buildXls(p)
		os.Remove(fileName)
	}
}
