package database

import "sync"

type Database struct {
	Collections map[string]Collection
	sync.RWMutex
}

func NewDatabase() *Database {
	var database Database
	database.Collections = make(map[string]Collection)
	return &database
}
