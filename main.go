package main

import (
	"SORA/Config"
	"SORA/Server"
)

func main() {
	Stop := make(chan int)
	if Config.StartGraphQLServer {
		Server.StartGraphQLServer()
	}
	<-Stop
}
