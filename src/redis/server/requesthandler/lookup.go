package requesthandler

import (
	"fmt"
	"redis/server/requesthandlers/hash"
)

var requestHandlers = map[string]RequestHandler{
	"SET": hash.Set,
	"GET": hash.Get,
}

func Lookup(command string) (RequestHandler, error) {

	requestHandler := requestHandlers[command]
	if requestHandler == nil {
		return nil, fmt.Errorf("Unknown command %s", command)
	}

	return requestHandler, nil
}
