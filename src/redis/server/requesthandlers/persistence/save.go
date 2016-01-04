package persistence

import (
	"fmt"
	"log"
	"os"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Save(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) > 0 {
		return fmt.Errorf("SAVE accepts not parameters")
	}

	return "OK"
}
