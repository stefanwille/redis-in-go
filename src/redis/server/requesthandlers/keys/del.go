package keys

import (
	"fmt"
	"redis/server/protocol"
	"redis/server/requesthandlers"
)

func Del(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) < 1 {
		return fmt.Errorf("DEL requires at least one KEY")
	}
	deleteCount := 0
	for _, key := range parameters {
		keyString, ok := key.(string)
		if !ok {
			return fmt.Errorf("DEL KEY must be a string")
		}
		_, found := requestContext.GetDatabase().Collections[keyString]
		if found {
			delete(requestContext.GetDatabase().Collections, keyString)
			deleteCount++
		}
	}

	return deleteCount
}
