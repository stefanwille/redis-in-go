package protocol

import (
	"bufio"
	"bytes"
	"testing"
)

func UnmarshalString(s string) (result Any, error error) {
	var buffer *bytes.Buffer = bytes.NewBufferString(s)
	reader := bufio.NewReader(buffer)
	return Unmarshal(reader)
}

func TestUnmarshalInt(t *testing.T) {
	result, error := UnmarshalString(":1000\r\n")
	if error != nil {
		t.Error(error)
	}
	if result.(int) != 1000 {
		t.Errorf("Expected 1000, got %v", result)
	}
}

func TestUnmarshalString(t *testing.T) {
	result, error := UnmarshalString("$6\r\nfoobar\r\n")
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "foobar" {
		t.Errorf("Expected foobar, got %v", result)
	}
}

func TestUnmarshalEmptyString(t *testing.T) {
	result, error := UnmarshalString("$0\r\n\r\n")
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "" {
		t.Errorf("Expected empty string, got %v", result)
	}
}

func TestUnmarshalNil(t *testing.T) {
	result, error := UnmarshalString("$-1\r\n")
	if error != nil {
		t.Error(error)
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestUnmarshalSimpleString(t *testing.T) {
	result, error := UnmarshalString("+foobar\r\n")
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "foobar" {
		t.Errorf("Expected foobar, got %v", result)
	}
}

func TestUnmarshalArray(t *testing.T) {
	result, error := UnmarshalString("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n")
	if error != nil {
		t.Error(error)
	}
	array := result.([]Any)
	if len(array) != 2 {
		t.Errorf("Expected array of length 2, got %v", array)
	}
}

func TestUnmarshalEmptyArray(t *testing.T) {
	result, error := UnmarshalString("*0\r\n")
	if error != nil {
		t.Error(error)
	}
	array := result.([]Any)
	if len(array) != 0 {
		t.Errorf("Expected array of length 0, got %v", array)
	}
}
