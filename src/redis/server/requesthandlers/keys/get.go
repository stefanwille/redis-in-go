package keys

import (
	"fmt"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Get(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) < 2 {
		return fmt.Errorf("GET requires at least KEY")
	}
	key, ok := request[1].(string)
	if !ok {
		return fmt.Errorf("GET KEY must be a string")
	}

	fmt.Printf("Collections: %v\n", requestContext.GetDatabase().Collections)

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
