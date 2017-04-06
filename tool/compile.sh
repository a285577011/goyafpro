#!/bin/bash

programDir="/Users/zenghui/go/mushu-youxing"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $programDir/bin/mushu-youxing-linux64 $programDir/src/main.go