package main

import (
	"game/gamenet"
	"log"
)

func main() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	gamenet.Start()
}
