package keys

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Keys(requestContext *requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 2 {
		return fmt.Errorf("KEYS requires at least a PATTERN")
	}
	_, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("KEYS PATTERN must be a string")
	}

	var keys []protocol.Any
	for key, _ := range (*requestContext).GetDatabase().Collections {
		keys = append(keys, key)
	}

	return keys
}
