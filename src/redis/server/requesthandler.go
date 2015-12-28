package server

import (
	"fmt"
	"redis/protocol"
)

type RequestHandler func(connection *Connection, request []protocol.Any) (response protocol.Any)

func set(connection *Connection, request []protocol.Any) (response protocol.Any) {
	if len(request) < 3 {
		return fmt.Errorf("SET requires at least KEY and VALUE")
	}
	key, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("SET KEY must be a string")
	}
	value, ok := request[2].(string)
	if !ok {
		return fmt.Errorf("SET VALUE must be a string")
	}

	collection := connection.database.collections[key]
	var hash *map[string]string
	if collection == nil {
		var newMap map[string]string
		newMap = make(map[string]string)
		hash = &newMap
		connection.database.collections[key] = hash
	} else {
		hash, ok = collection.(*map[string]string)
		if !ok {
			return fmt.Errorf("KEY is not a HASH, but a %T", collection)
		}
	}

	(*hash)[key] = value

	return "OK"
}

func get(connection *Connection, request []protocol.Any) (response protocol.Any) {
	if len(request) < 2 {
		return fmt.Errorf("GET requires at least KEY")
	}
	key, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("GET KEY must be a string")
	}

	collection := connection.database.collections[key]
	var hash *map[string]string
	if collection == nil {
		return nil
	} else {
		hash, ok = collection.(*map[string]string)
		if !ok {
			return fmt.Errorf("KEY is not a HASH, but a %T", collection)
		}
	}

	return (*hash)[key]
}
