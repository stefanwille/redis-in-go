package keys

import (
	"redis/server/protocol"
	"redis/server/requesthandlers"
	"testing"
)

func TestDel(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key1"] = "value"
	response := Del(requestContext, []protocol.Any{"key1", "key2"})
	if response != 1 {
		t.Errorf("Expected response 1, got %v", response)
		return
	}
	key1 := requestContext.GetDatabase().Collections["key1"]
	if key1 != nil {
		t.Errorf("Expected value key1 to equal nil, got %v", key1)
	}
}
