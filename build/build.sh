#!/bin/sh

ver=v0.2.0

env GOOS=linux go build -o "cvs2xlsx.linux.${ver}" ../
env GOOS=darwin go build -o "cvs2xlsx.mac.${ver}" ../
env GOOS=windows go build -o "cvs2xlsx.win.${ver}.exe" ../