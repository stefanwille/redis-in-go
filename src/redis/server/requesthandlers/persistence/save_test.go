package persistence

import (
	"redis/server/protocol"
	"redis/server/requesthandlers"
	"testing"
)

func TestSave_WritesAFile(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key"] = "value"

	response := Save(requestContext, []protocol.Any{})
	if response != "OK" {
		t.Errorf("Expected response 'OK', got %v", response)
		return
	}
}
