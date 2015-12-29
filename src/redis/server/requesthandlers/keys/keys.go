package keys

import (
	"fmt"
	"path/filepath"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Keys(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 1 {
		return fmt.Errorf("KEYS requires at least a PATTERN")
	}
	pattern, ok := request[0].(string)
	if !ok {
		return fmt.Errorf("KEYS PATTERN must be a string")
	}

	var keys []protocol.Any
	for key, _ := range requestContext.GetDatabase().Collections {
		match, error := filepath.Match(pattern, key)
		if error != nil {
			return error
		}
		if match {
			keys = append(keys, key)
		}
	}

	return keys
}
