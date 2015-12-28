package server

type Database struct {
	objects map[string]Object
}

func NewDatabase() *Database {
	var database Database
	return &database
}
