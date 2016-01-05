package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"redis/server/database"
	"redis/server/protocol"
	"redis/server/requesthandler"
	"runtime/debug"
	"strings"
)

type Connection struct {
	conn     net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
	database *database.Database
}

func NewConnection(conn net.Conn, database *database.Database) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	connection := Connection{conn, reader, writer, database}
	return &connection
}

func (connection *Connection) GetDatabase() *database.Database {
	return connection.database
}

func (connection *Connection) ServeRequests() {
	log.Printf("New connection from %v\n", connection.conn.RemoteAddr())
	defer func() {
		if r := recover(); r != nil {

			log.Printf("*** Panic ***: %v, %s", r, debug.Stack())
			return
		}
	}()
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
	requestSlice, ok := request.([]protocol.Any)
	if !ok {
		return fmt.Errorf("Expected request to be an array, got %T", request)
	}
	command, ok := requestSlice[0].(string)
	if !ok {
		return fmt.Errorf("Expected command to be a string, got %T", command)
	}
	log.Printf("Request: %v", requestSlice)
	command = strings.ToUpper(command)
	requestHandler, isWriter, error := requesthandler.Lookup(command)
	if error != nil {
		return error
	}

	if isWriter {
		connection.database.Lock()
		defer connection.database.Unlock()
	} else {
		connection.database.RLock()
		defer connection.database.RUnlock()

	}

	return requestHandler(connection, requestSlice[1:])
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
