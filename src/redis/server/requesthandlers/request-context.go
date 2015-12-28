package requesthandlers

import "redis/server/database"

type RequestContext interface {
	GetDatabase() *database.Database
}
