package requesthandler

import (
	"redis/server/protocol"
	"redis/server/requesthandlers"
)

type RequestHandler func(requestContext requesthandlers.RequestContext, request []protocol.Any) protocol.Any
