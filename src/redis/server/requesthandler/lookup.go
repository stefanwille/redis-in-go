package requesthandler

import (
	"fmt"
	"redis/protocol"
)

func Lookup(command string) (RequestHandler, error) {
	switch command {
	case "SET":
		return set, nil
	case "GET":
		return get, nil
	default:
		return nil, fmt.Errorf("Unknown command %s", command)
	}
}

func set(requestContext *RequestContext, request []protocol.Any) (response protocol.Any) {
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

	collection := (*requestContext).GetDatabase().Collections[key]
	var hash *map[string]string
	if collection == nil {
		var newMap map[string]string
		newMap = make(map[string]string)
		hash = &newMap
		(*requestContext).GetDatabase().Collections[key] = hash
	} else {
		hash, ok = collection.(*map[string]string)
		if !ok {
			return fmt.Errorf("KEY is not a HASH, but a %T", collection)
		}
	}

	(*hash)[key] = value

	return "OK"
}

func get(requestContext *RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 2 {
		return fmt.Errorf("GET requires at least a KEY")
	}
	key, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("GET KEY must be a string")
	}

	collection := (*requestContext).GetDatabase().Collections[key]
	if collection == nil {
		return nil
	}

	hash, ok := collection.(*map[string]string)
	if !ok {
		return fmt.Errorf("KEY is not a HASH, but a %T", collection)
	}

	return (*hash)[key]
}
