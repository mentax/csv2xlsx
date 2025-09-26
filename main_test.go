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

func TestBuildXls_WithMaxMemory(t *testing.T) {
	dummyCSV := "./dummy_large.csv"
	outputFile := "./test_large_output.xlsx"
	defer os.Remove(dummyCSV)
	defer os.Remove(outputFile)

	// Create a large dummy CSV file to test disk-based storage.
	file, err := os.Create(dummyCSV)
	assert.Nil(t, err)
	for i := 0; i < 50000; i++ {
		_, err := file.WriteString("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t\n")
		assert.Nil(t, err)
	}
	file.Close()

	params := &params{
		input:     []string{dummyCSV},
		output:    outputFile,
		maxMemory: true, // Enable disk-based storage
		delimiter: ',',
	}

	err = buildXls(params)
	assert.Nil(t, err, "buildXls should complete without error when using disk-based storage")
	assert.FileExists(t, outputFile)
}
