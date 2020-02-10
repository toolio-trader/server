#!/usr/bin/env bash

CGO_ENABLED=0
GOOS=linux
GOARCH=darwin

mkdir -p ./build/$GOOS/$GOARCH
GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o ./build/$GOOS/$GOARCH

CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

mkdir -p ./build/$GOOS/$GOARCH
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/$GOOS/$GOARCH

GOOS=windows
GOARCH=386

mkdir -p ./build/$GOOS/$GOARCH
GOOS=windows GOARCH=386 go  build -ldflags="-w -s" -o ./build/$GOOS/$GOARCH