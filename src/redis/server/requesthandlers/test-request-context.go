package requesthandlers

import (
	"redis/server/database"
)

type TestRequestContext struct {
	database *database.Database
}

func NewTestRequestContext() *TestRequestContext {
	var testRequestContext TestRequestContext
	testRequestContext.database = database.NewDatabase()
	return &testRequestContext
}

func (testRequestContext *TestRequestContext) GetDatabase() *database.Database {
	return testRequestContext.database
}
