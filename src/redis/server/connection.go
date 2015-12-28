package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"redis/protocol"
	"strings"
)

type Connection struct {
	conn     net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
	database *Database
}

func NewConnection(conn net.Conn, database *Database) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	connection := Connection{conn, reader, writer, database}
	return &connection
}

func (connection *Connection) ServeRequests() {
	log.Printf("New connection from %v\n", connection.conn.RemoteAddr())
	defer connection.conn.Close()
	for {
		request, eof, error := connection.receiveRequest()
		if error != nil {
			continue
		}
		if eof {
			return
		}

		response := connection.handleRequest(request)

		error = connection.sendResponse(response)
		if error != nil {
			return
		}
	}
}

func (connection *Connection) receiveRequest() (request protocol.Any, eof bool, error error) {
	request, error = protocol.Unmarshal(connection.reader)
	if error != nil {
		log.Println(error)
		if error.Error() == "EOF" {
			return nil, true, nil
		}

		error := connection.sendErrorResponse(error)
		return nil, false, error
	}

	return request, false, nil
}

func (connection *Connection) handleRequest(request protocol.Any) (response protocol.Any) {
	log.Printf("handleRequest %v", request)
	requestSlice, ok := request.([]protocol.Any)
	if !ok {
		return fmt.Errorf("Expected request to be an array, got %T", request)
	}
	command, ok := requestSlice[0].(string)
	if !ok {
		return fmt.Errorf("Expected command to be a string, got %T", command)
	}
	command = strings.ToUpper(command)
	switch command {
	case "SET":
		return connection.set(requestSlice)
	case "GET":
		return connection.get(requestSlice)
	default:
		return fmt.Errorf("Unknown command %s", command)
	}
}

func (connection *Connection) sendErrorResponse(response error) (error error) {
	error = connection.sendResponse(response)
	if error != nil {
		return error
	}

	return nil
}

func (connection *Connection) sendResponse(response protocol.Any) (error error) {
	error = protocol.Marshal(connection.writer, response)
	if error != nil {
		log.Printf("Error while sending response: %v", error)
		return error
	}

	error = connection.writer.Flush()
	if error != nil {
		log.Printf("Error while flushing response: %v", error)
		return error
	}

	return nil
}

func (connection *Connection) set(request []protocol.Any) (error error) {
	return nil
}

func (connection *Connection) get(request []protocol.Any) (error error) {
	return nil
}
