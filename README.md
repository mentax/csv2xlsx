

[![Build Status](https://travis-ci.org/mentax/cvs2xlsx.svg?branch=master)](https://travis-ci.org/mentax/cvs2xlsx)
[![GoDoc](https://godoc.org/github.com/mentax/cvs2xlsx?status.svg)](https://godoc.org/github.com/mentax/cvs2xlsx)
[![codebeat badge](https://codebeat.co/badges/042f1764-a799-4a7d-abd3-80664e7ce257)](https://codebeat.co/projects/github-com-mentax-cvs2xlsx-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mentax/cvs2xlsx)](https://goreportcard.com/report/github.com/mentax/cvs2xlsx)
[![Coverage](https://gocover.io/_badge/github.com/mentax/cvs2xlsx)](http://gocover.io/github.com/mentax/cvs2xlsx)

# cvs 2 xlsx

## HELP 
  Actual version always on  cvs2xlsx -h or cvs2xlsx help

### NAME:
   cvs2xlsx - Convert CSV data to xlsx - especially the big one. 
   
### Speed: 
   
   cvs with 50k rows, 5 MB, with xlsx template - 5s
   
   
   (On MacBook Pro 2016) 

### Example: 

```bash
cvs2xlsx --template example/template.xlsx --sheet Sheet_1 --sheet Sheet_2 --row 2 --output result.xlsx data.csv data2.csv 
cvs2xlsx.exe -t example  emplate.xlsx -s Sheet_1 -s Sheet_2 -r 2 -o result.xlsx data.csv data2.csv 
```

### USAGE:

    cvs2xlsx [global options] command [command options] [file of file's list with csv data]

#### VERSION:
   0.2.0
  
#### GLOBAL OPTIONS:

```
--sheets names, -s names          sheet names in the same order like cvs files. If sheet with that name exists, data is inserted to this sheet. Usage: -s AA -s BB
--template path, -t path          path to xlsx file with template output
--row number, -r number           row number to use for create rows format. When '0' - not used. This row will be removed from xlsx file. (default: 0)
--output xlsx file, -o xlsx file  path to result xlsx file (default: "./output.xlsx")
--help, -h                        show help
--version, -v                     print the version
```   
   
   
## Download

Download from [releases section on GitHub](https://github.com/mentax/cvs2xlsx/releases)   