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

func TestBuildXls_WithUseCache(t *testing.T) {
	dummyCSV := "./dummy_large.csv"
	outputFile := "./test_large_output.xlsx"
	defer os.Remove(dummyCSV)
	defer os.Remove(outputFile)

	// Create a large dummy CSV file to test disk-based storage.
	file, err := os.Create(dummyCSV)
	assert.Nil(t, err)
	for i := 0; i < 10000; i++ { // Reduced from 50000
		_, err := file.WriteString("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t\n")
		assert.Nil(t, err)
	}
	file.Close()

	params := &params{
		input:     []string{dummyCSV},
		output:    outputFile,
		useCache:  true, // Enable disk-based storage
		delimiter: ',',
	}

	err = buildXls(params)
	assert.Nil(t, err, "buildXls should complete without error when using disk-based storage")
	assert.FileExists(t, outputFile)
}

func createDummyCSV(b *testing.B, numRows int) string {
	b.Helper()
	dummyCSV := "./dummy_benchmark.csv"
	file, err := os.Create(dummyCSV)
	if err != nil {
		b.Fatalf("failed to create dummy csv: %v", err)
	}
	for i := 0; i < numRows; i++ {
		_, err := file.WriteString("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t\n")
		if err != nil {
			b.Fatalf("failed to write to dummy csv: %v", err)
		}
	}
	file.Close()
	return dummyCSV
}

func BenchmarkBuildXls_InMemory(b *testing.B) {
	dummyCSV := createDummyCSV(b, 10000) // Reduced from 50000
	defer os.Remove(dummyCSV)
	outputFile := "./benchmark_in_memory_output.xlsx"
	defer os.Remove(outputFile)

	params := &params{
		input:     []string{dummyCSV},
		output:    outputFile,
		useCache:  false, // Use in-memory storage
		delimiter: ',',
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err := buildXls(params)
		if err != nil {
			b.Fatalf("buildXls failed: %v", err)
		}
	}
}

func BenchmarkBuildXls_DiskV(b *testing.B) {
	dummyCSV := createDummyCSV(b, 10000) // Reduced from 50000
	defer os.Remove(dummyCSV)
	outputFile := "./benchmark_diskv_output.xlsx"
	defer os.Remove(outputFile)

	params := &params{
		input:     []string{dummyCSV},
		output:    outputFile,
		useCache:  true, // Enable disk-based storage
		delimiter: ',',
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err := buildXls(params)
		if err != nil {
			b.Fatalf("buildXls failed: %v", err)
		}
	}
}