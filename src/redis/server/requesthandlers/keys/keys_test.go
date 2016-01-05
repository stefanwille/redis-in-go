package keys

import (
	"redis/server/protocol"
	"redis/server/requesthandlers"
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeys(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key"] = "value"

	response := Keys(requestContext, []protocol.Any{"*"})
	if !reflect.DeepEqual(response, []protocol.Any{"key"}) {
		t.Errorf("Expected response ['key'], got %v", response)
		return
	}
}

func TestKeys_FiltersByGlob(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key"] = "value"

	response := Keys(requestContext, []protocol.Any{"m*"})
	sa, ok := response.([]protocol.Any)
	if !ok {
		t.Errorf("Expected response [], got %T ", response)
	}

	if len(sa) != 0 {
		t.Errorf("Expected response [], got %v %d", sa, len(sa))
		return
	}
}
