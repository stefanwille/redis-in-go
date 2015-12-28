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
		request, eof, error := receiveRequest(reader, writer)
		if error != nil {
			continue
		}
		if eof {
			return
		}

		response, error := handleRequest(request)
		if error != nil {
			return
		}

		error = sendResponse(writer, response)
		if error != nil {
			return
		}
	}
}

func receiveRequest(reader *bufio.Reader, writer *bufio.Writer) (request protocol.Any, eof bool, error error) {
	request, error = protocol.Unmarshal(reader)
	if error != nil {
		log.Println(error)
		if error.Error() == "EOF" {
			return nil, true, nil
		}

		error := sendErrorResponse(writer, error)
		return nil, false, error
	}

	return request, false, nil
}

func handleRequest(request protocol.Any) (response protocol.Any, error error) {
	log.Printf("handleRequest %v", request)
	return "OK", nil
}

func sendErrorResponse(writer *bufio.Writer, response error) (error error) {
	error = sendResponse(writer, response)
	if error != nil {
		return error
	}

	return nil
}

func sendResponse(writer *bufio.Writer, response protocol.Any) (error error) {
	error = protocol.Marshal(writer, response)
	if error != nil {
		log.Printf("Error while sending response: %v", error)
		return error
	}

	error = writer.Flush()
	if error != nil {
		log.Printf("Error while flushing response: %v", error)
		return error
	}

	return nil
}
