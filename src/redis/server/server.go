package server

import (
	"log"
	"net"
	"redis/server/database"
)

type Server struct {
	listenAddress string
	database      *database.Database
}

func New(listenAddress string) *Server {
	database := database.NewDatabase()
	server := Server{listenAddress, database}
	return &server
}

func (server *Server) Listen() (error error) {
	log.Printf("Launching server on address %s", server.listenAddress)
	listener, error := net.Listen("tcp", server.listenAddress)
	if error != nil {
		log.Fatal(error)
		return error
	}
	defer listener.Close()

	for {
		connection, error := listener.Accept()
		if error != nil {
			log.Print(error)
			return error
		}
		go server.handleConnection(connection)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	connection := NewConnection(conn, server.database)
	connection.ServeRequests()
}
