package keys

import (
	"fmt"
	"redis/server/protocol"
	"redis/server/requesthandlers"
)

func Set(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) < 2 {
		return fmt.Errorf("SET requires at least KEY and VALUE")
	}
	key, ok := parameters[0].(string)
	if !ok {
		return fmt.Errorf("SET KEY must be a string")
	}
	value, ok := parameters[1].(string)
	if !ok {
		return fmt.Errorf("SET VALUE must be a string")
	}

	requestContext.GetDatabase().Collections[key] = value

	return "OK"
}
