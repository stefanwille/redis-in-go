package requesthandler

import (
	"fmt"
	"redis/server/requesthandlers/general"
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
	{command: "DBSIZE", requestHandler: general.Dbsize, isWriter: false},
}

var commandMap = map[string]*RequestHandlerDefinition{}

func init() {
	buildCommandMap()
}

func buildCommandMap() {
	for _, requestHandlerDefinition := range requestHandlerDefinitions {
		commandMap[requestHandlerDefinition.command] = requestHandlerDefinition
	}
}

func Lookup(command string) (RequestHandler, bool, error) {
	var requestHandlerDefinition *RequestHandlerDefinition = commandMap[command]
	if requestHandlerDefinition == nil {
		return nil, true, fmt.Errorf("Unknown command %s", command)
	}

	return requestHandlerDefinition.requestHandler, requestHandlerDefinition.isWriter, nil
}
