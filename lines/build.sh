#!/bin/zsh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./built/lines_linux .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./built/lines.exe .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./built/lines_darwin .