#/bin/bash
CURR=$(pwd)
GOPATH=$GOPATH:$CURR
gofmt -w ./src/game

go build -o ./bin/server ./src/game
