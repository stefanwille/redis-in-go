package keys

import (
	"redis/protocol"
	"redis/server/requesthandlers"
	"testing"
)

func TestGet_ReturnsStringValue(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key"] = "value"

	response := Get(requestContext, []protocol.Any{"GET", "key"})
	if response != "value" {
		t.Errorf("Expected response 'value', got %v", response)
		return
	}
}

func TestGet_ReturnsNil(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()
	requestContext.GetDatabase().Collections["key"] = nil

	response := Get(requestContext, []protocol.Any{"GET", "key"})
	if response != nil {
		t.Errorf("Expected response nil, got %v", response)
		return
	}
}

func TestGet_ReturnsUndefinedAsNil(t *testing.T) {
	requestContext := requesthandlers.NewTestRequestContext()

	response := Get(requestContext, []protocol.Any{"GET", "key"})
	if response != nil {
		t.Errorf("Expected response nil, got %v", response)
		return
	}
}
