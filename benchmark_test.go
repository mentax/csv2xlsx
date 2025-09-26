//go:build benchmark
// +build benchmark

package main

import (
	"fmt"
	"os"
	"testing"
)

const BENCHMARK_FILE_NAME = "benchmark.csv"
const BENCHMARK_ROWS = 50000

func createBenchmarkFile() {
	f, err := os.Create(BENCHMARK_FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < BENCHMARK_ROWS; i++ {
		f.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", i, i+1, i+2, i+3))
	}
}

func removeBenchmarkFile() {
	os.Remove(BENCHMARK_FILE_NAME)
}

func BenchmarkBuildXlsWithBigFile(b *testing.B) {
	createBenchmarkFile()
	defer removeBenchmarkFile()

	p := &params{
		input:  []string{BENCHMARK_FILE_NAME},
		output: "./testFile.xlsx",
	}

	for n := 0; n < b.N; n++ {
		buildXls(p)
		os.Remove(p.output)
	}
}