package keys

import (
	"redis/server/protocol"
	"redis/server/requesthandlers"
	"testing"
)

func TestSet(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	response := Set(requestContext, []protocol.Any{"key", "value"})
	if response != "OK" {
		t.Errorf("Expected response OK, got %v", response)
		return
	}
	key := requestContext.GetDatabase().Collections["key"]
	if key != "value" {
		t.Errorf("Expected value key to equal 'value', got %v", key)
	}
}
