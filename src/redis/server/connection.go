package server

import (
	"bufio"
	"log"
	"net"
	"redis/protocol"
)

type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewConnection(conn net.Conn) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	connection := Connection{conn, reader, writer}
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

		response, error := connection.handleRequest(request)
		if error != nil {
			return
		}

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

func (connection *Connection) handleRequest(request protocol.Any) (response protocol.Any, error error) {
	log.Printf("handleRequest %v", request)
	return "OK", nil
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
