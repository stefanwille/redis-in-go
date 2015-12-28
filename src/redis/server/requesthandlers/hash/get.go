package hash

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

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
