package hash

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Set(requestContext *requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
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

func Get(requestContext *requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
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
