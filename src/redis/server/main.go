package main

import (
	"bufio"
	"log"
	"net"
	"redis/protocol"
)

const listenAddress = ":6379"

func main() {
	log.Printf("Launching server on address %s", listenAddress)
	listener, error := net.Listen("tcp", listenAddress)
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
	defer listener.Close()

	for {
		connection, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
			panic(error)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	log.Printf("New connection from %v\n", connection.RemoteAddr())
	defer connection.Close()
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)
	for {
		request, error := protocol.Unmarshal(reader)
		if error != nil {
			log.Println(error)
			if error.Error() == "EOF" {
				return
			}
		}

		log.Printf("request %v", request)
		protocol.Marshal(writer, "OK")
		writer.Flush()
	}
}
