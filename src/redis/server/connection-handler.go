package server

import (
	"bufio"
	"log"
	"net"
	"redis/protocol"
)

type ConnectionHandler struct {
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func NewConnectionHandler(connection net.Conn) *ConnectionHandler {
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)
	connectionHandler := ConnectionHandler{connection, reader, writer}
	return &connectionHandler
}

func (connectionHandler *ConnectionHandler) handleConnection() {
	log.Printf("New connection from %v\n", connectionHandler.connection.RemoteAddr())
	defer connectionHandler.connection.Close()
	for {
		request, eof, error := connectionHandler.receiveRequest()
		if error != nil {
			continue
		}
		if eof {
			return
		}

		response, error := connectionHandler.handleRequest(request)
		if error != nil {
			return
		}

		error = connectionHandler.sendResponse(response)
		if error != nil {
			return
		}
	}
}

func (connectionHandler *ConnectionHandler) receiveRequest() (request protocol.Any, eof bool, error error) {
	request, error = protocol.Unmarshal(connectionHandler.reader)
	if error != nil {
		log.Println(error)
		if error.Error() == "EOF" {
			return nil, true, nil
		}

		error := connectionHandler.sendErrorResponse(error)
		return nil, false, error
	}

	return request, false, nil
}

func (connectionHandler *ConnectionHandler) handleRequest(request protocol.Any) (response protocol.Any, error error) {
	log.Printf("handleRequest %v", request)
	return "OK", nil
}

func (connectionHandler *ConnectionHandler) sendErrorResponse(response error) (error error) {
	error = connectionHandler.sendResponse(response)
	if error != nil {
		return error
	}

	return nil
}

func (connectionHandler *ConnectionHandler) sendResponse(response protocol.Any) (error error) {
	error = protocol.Marshal(connectionHandler.writer, response)
	if error != nil {
		log.Printf("Error while sending response: %v", error)
		return error
	}

	error = connectionHandler.writer.Flush()
	if error != nil {
		log.Printf("Error while flushing response: %v", error)
		return error
	}

	return nil
}
