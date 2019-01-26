package main

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Config"
	"SORA/Project/Go_Back_End_SEGA_Project/Server"
)

func main() {
	Stop := make(chan int)
	if Config.StartGraphQLServer {
		Server.StartGraphQLServer()
	}
	<-Stop
}
