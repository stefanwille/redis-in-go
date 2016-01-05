package persistence

import (
	"fmt"
	"redis/server/protocol"
	"redis/server/requesthandlers"
)

func Save(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) > 0 {
		return fmt.Errorf("SAVE accepts no parameters")
	}

	requestContext.GetDatabase().Save("/tmp/database.json")

	return "OK"
}
