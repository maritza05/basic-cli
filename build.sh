#!/usr/bin/env sh

GOOS=linux GOARCH=amd64 go build -v -o pricing-engine-linux main.go
GOOS=darwin GOARCH=amd64 go build -v -o pricing-engine-osx main.go
