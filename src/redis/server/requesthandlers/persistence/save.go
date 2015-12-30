package persistence

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
	"log"
	"redis/protocol"
	"redis/server/requesthandlers"
)

func Save(requestContext requesthandlers.RequestContext, request []protocol.Any) (response protocol.Any) {
	if len(request) > 0 {
		return fmt.Errorf("SAVE accepts not parameters")
	}

	// collections := requestContext.GetDatabase().Collections

	sqlite, error := sqlite3.Open("sqlite.db")
	if error != nil {
		log.Print(error)
		return error
	}
	defer sqlite.Close()

	sqlite.Exec("CREATE TABLE x(a, b, c)")
	sqlite.Commit()

	return "OK"
}
