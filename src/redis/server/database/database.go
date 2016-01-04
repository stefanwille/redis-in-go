package database

import "sync"
import "os"
import "log"
import "encoding/json"

type Database struct {
	Collections  map[string]Collection `json:"Objects"`
	sync.RWMutex `json:"-"`
}

func NewDatabase() *Database {
	var database Database
	database.Collections = make(map[string]Collection)
	return &database
}

func (database *Database) Load(filename string) error {
	log.Printf("Loading database %s", filename)
	file, error := os.Open(filename)
	if error != nil {
		return error
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(database)

	return nil
}

func (database *Database) Save(filename string) error {
	file, error := os.Create("/tmp/database.json")
	if error != nil {
		log.Printf("Error while creating database file: %v", error)
		return error
	}

	defer file.Close()
	log.Printf("Saving database at %s", filename)
	encoder := json.NewEncoder(file)
	error = encoder.Encode(database)
	if error != nil {
		log.Printf("Error while save database: %v", error)
		return error
	}
	log.Printf("Done saving database")
	return nil
}
