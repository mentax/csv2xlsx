#!/bin/sh

ver=v0.2.2

env GOOS=linux   go build -o "csv2xlsx.linux.${ver}"   ../
env GOOS=darwin  go build -o "csv2xlsx.mac.${ver}"     ../
env GOOS=windows go build -o "csv2xlsx.win.${ver}.exe" ../
