package keys

import (
	"redis/protocol"
	"redis/server/requesthandlers"
	"testing"
)

func TestSet(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	response := Set(requestContext, []protocol.Any{"foo", "bar"})
	if response != "OK" {
		t.Errorf("Expected response OK, got %v", response)
		return
	}
	foo := requestContext.GetDatabase().Collections["foo"].(string)
	if foo != "bar" {
		t.Errorf("Expected value foo to equal bar, got %v", foo)
	}
}
