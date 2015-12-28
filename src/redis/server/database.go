package server

import "sync"

type Database struct {
	collections map[string]Collection
	sync.Mutex
}

func NewDatabase() *Database {
	var database Database
	database.collections = make(map[string]Collection)
	return &database
}
