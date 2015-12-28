package requesthandler

import (
	"redis/protocol"
)

type RequestHandler func(requestContext *RequestContext, request []protocol.Any) (response protocol.Any)
