package main

import "redis/server"

const listenAddress = ":6379"

func main() {
	server := server.New(listenAddress)
	error := server.Load("/tmp/database.json")
	if error != nil {
		panic(error)
	}
	server.Listen()
}
