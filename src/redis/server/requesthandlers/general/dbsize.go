package general

import (
	"fmt"
	"redis/server/protocol"
	"redis/server/requesthandlers"
)

func Dbsize(requestContext requesthandlers.RequestContext, parameters []protocol.Any) (response protocol.Any) {
	if len(parameters) > 0 {
		return fmt.Errorf("DBSIZE accepts no parameters")
	}

	panic("This is a test")

	return "OK"
}
