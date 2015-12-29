package keys

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Set(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 2 {
		return fmt.Errorf("SET requires at least KEY and VALUE")
	}
	key, ok := request[0].(string)
	if !ok {
		return fmt.Errorf("SET KEY must be a string")
	}
	value, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("SET VALUE must be a string")
	}

	requestContext.GetDatabase().Collections[key] = value

	return "OK"
}
