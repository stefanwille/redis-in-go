package requesthandler

import (
	"fmt"
	// "redis/server/requesthandlers/hash"
	"redis/server/requesthandlers/keys"
)

type RequestHandlerDefinition struct {
	command        string
	requestHandler RequestHandler
	isWriter       bool
}

var requestHandlerDefinitions = []RequestHandlerDefinition{
	// "HSET": hash.Hset,
	// "HGET": hash.Hget,
	{command: "GET", requestHandler: keys.Get, isWriter: false},
	{command: "SET", requestHandler: keys.Set, isWriter: true},
	{command: "DEL", requestHandler: keys.Del, isWriter: true},
	{command: "KEYS", requestHandler: keys.Keys, isWriter: false},
}

var commandToRequestHandlerDefinition = map[string]*RequestHandlerDefinition{}

func init() {
	for _, requestHandlerDefinition := range requestHandlerDefinitions {
		commandToRequestHandlerDefinition[requestHandlerDefinition.command] = &requestHandlerDefinition
	}
}

func Lookup(command string) (RequestHandler, bool, error) {
	var requestHandlerDefinition *RequestHandlerDefinition = commandToRequestHandlerDefinition[command]
	if requestHandlerDefinition == nil {
		return nil, true, fmt.Errorf("Unknown command %s", command)
	}

	return requestHandlerDefinition.requestHandler, requestHandlerDefinition.isWriter, nil
}
