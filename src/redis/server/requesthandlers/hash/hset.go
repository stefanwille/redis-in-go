package hash

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Hset(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) < 3 {
		return fmt.Errorf("SET requires at least KEY and VALUE")
	}
	key, ok := parameters[1].(string)
	if !ok {
		return fmt.Errorf("SET KEY must be a string")
	}
	value, ok := parameters[2].(string)
	if !ok {
		return fmt.Errorf("SET VALUE must be a string")
	}

	collection := requestContext.GetDatabase().Collections[key]
	var hash *map[string]string
	if collection == nil {
		var newMap map[string]string
		newMap = make(map[string]string)
		hash = &newMap
		requestContext.GetDatabase().Collections[key] = hash
	} else {
		hash, ok = collection.(*map[string]string)
		if !ok {
			return fmt.Errorf("KEY is not a HASH, but a %T", collection)
		}
	}

	(*hash)[key] = value

	return "OK"
}
