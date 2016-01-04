package keys

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Get(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) < 1 {
		return fmt.Errorf("GET requires at least KEY")
	}
	key, ok := parameters[0].(string)
	if !ok {
		return fmt.Errorf("GET KEY must be a string")
	}

	value := requestContext.GetDatabase().Collections[key]

	if value == nil {
		return value
	}

	stringValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("GET value must be a string, but is a %T", value)
	}

	return stringValue
}
