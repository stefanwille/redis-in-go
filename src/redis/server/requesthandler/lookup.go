package requesthandler

import (
	"fmt"
	"redis/server/requesthandlers/hash"
	"redis/server/requesthandlers/keys"
)

var requestHandlers = map[string]RequestHandler{
	"SET":  hash.Set,
	"GET":  hash.Get,
	"KEYS": keys.Keys,
}

func Lookup(command string) (RequestHandler, error) {

	requestHandler := requestHandlers[command]
	if requestHandler == nil {
		return nil, fmt.Errorf("Unknown command %s", command)
	}

	return requestHandler, nil
}
