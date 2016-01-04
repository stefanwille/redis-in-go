package requesthandler

import (
	"fmt"
	"redis/server/requesthandlers/keys"
	"redis/server/requesthandlers/persistence"
)

type RequestHandlerDefinition struct {
	command        string
	requestHandler RequestHandler
	isWriter       bool
}

var requestHandlerDefinitions = []*RequestHandlerDefinition{
	{command: "GET", requestHandler: keys.Get, isWriter: false},
	{command: "SET", requestHandler: keys.Set, isWriter: true},
	{command: "DEL", requestHandler: keys.Del, isWriter: true},
	{command: "KEYS", requestHandler: keys.Keys, isWriter: false},
	{command: "SAVE", requestHandler: persistence.Save, isWriter: false},
}

var commandToRequestHandlerDefinition = map[string]*RequestHandlerDefinition{}

func init() {
	for _, requestHandlerDefinition := range requestHandlerDefinitions {
		commandToRequestHandlerDefinition[requestHandlerDefinition.command] = requestHandlerDefinition
	}
}

func Lookup(command string) (RequestHandler, bool, error) {
	var requestHandlerDefinition *RequestHandlerDefinition = commandToRequestHandlerDefinition[command]
	if requestHandlerDefinition == nil {
		return nil, true, fmt.Errorf("Unknown command %s", command)
	}

	return requestHandlerDefinition.requestHandler, requestHandlerDefinition.isWriter, nil
}
