package keys

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Delete(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 1 {
		return fmt.Errorf("DEL requires at least one KEY")
	}
	deleteCount := 0
	for _, key := range request {
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
