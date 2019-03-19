package main

import (
	"SORA/Command"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	Command.Start()
}
