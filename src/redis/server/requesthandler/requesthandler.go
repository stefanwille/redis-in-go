package requesthandler

import (
	"redis/protocol"
	"redis/server/requesthandlers"
)

type RequestHandler func(requestContext requesthandlers.RequestContext, request []protocol.Any) protocol.Any
