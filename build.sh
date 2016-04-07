#/bin/bash

gofmt -w ./src
go build -o ./bin/server ./src
