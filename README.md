

[![goreleaser](https://github.com/mentax/csv2xlsx/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/mentax/csv2xlsx/actions/workflows/goreleaser.yml)
[![test](https://github.com/mentax/csv2xlsx/actions/workflows/test.yaml/badge.svg)](https://github.com/mentax/csv2xlsx/actions/workflows/test.yaml)
[![GoDoc](https://godoc.org/github.com/mentax/csv2xlsx?status.svg)](https://godoc.org/github.com/mentax/csv2xlsx)
[![codebeat badge](https://codebeat.co/badges/1b57272c-e0fa-4a14-93b5-3586e192fdb3)](https://codebeat.co/projects/github-com-mentax-csv2xlsx-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mentax/csv2xlsx)](https://goreportcard.com/report/github.com/mentax/csv2xlsx)
[![GitHub release](https://img.shields.io/github/v/release/mentax/csv2xlsx)](https://github.com/mentax/csv2xlsx/releases/latest)

<!-- 
  [![Coverage](https://gocover.io/_badge/github.com/mentax/csv2xlsx)](http://gocover.io/github.com/mentax/csv2xlsx)
-->

# csv 2 xlsx

## HELP
  An actual version always available by run `csv2xlsx -h` or `csv2xlsx help`

### NAME:
   csv2xlsx - Convert CSV data to xlsx - especially the big one.

### Speed:

   csv with 50k rows, 5 MB, with xlsx template - 5s


   (On MacBook Pro 2016)

### Example:

```bash
csv2xlsx --template example/template.xlsx --sheet Sheet_1 --sheet Sheet_2 --row 2 --output result.xlsx data.csv data2.csv
csv2xlsx.exe -t example/template.xlsx -s Sheet_1 -s Sheet_2 -r 2 -o result.xlsx data.csv data2.csv
```

### USAGE:

    csv2xlsx [global options] command [command options] [file of file's list with csv data]

#### GLOBAL OPTIONS:

```
 --sheets names, -s names [ --sheets names, -s names ]  sheet names in the same order like csv files. If sheet with that name exists, data is inserted to this sheet. Usage: -s AA -s BB
 --template path, -t path                               path to xlsx file with template file
 --delimiter letter, -d letter                          one letter delimiter used in csv file (default: ",")
 --exampleRow number, -r number                         exampleRow number to use for create rows format. When '0' - not used. This exampleRow will be overwrite in result file. (default: 0)
 --startFrom number, --sf number                        startFrom number decide which row is used as first row from csv file. Counting from 0. (default: 0)
 --output xlsx file, -o xlsx file                       path to result xlsx file (default: "./output.xlsx")
 --help, -h                                             show help
 --version, -v                                          print the version
```   


## Download

Download from [releases section on GitHub](https://github.com/mentax/csv2xlsx/releases)   
