package main

import "redis/server"

const listenAddress = ":6379"

func main() {
	server := server.New(listenAddress)
	server.Listen()
}
