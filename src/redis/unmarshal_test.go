package redis

import (
	"bytes"
	"testing"
)

func TestUnmarshalInt(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString(":1000\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	if result.(int) != 1000 {
		t.Errorf("Expected 1000, got %v", result)
	}
}

func TestUnmarshalString(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString("$6\r\nfoobar\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "foobar" {
		t.Errorf("Expected foobar, got %v", result)
	}
}

func TestUnmarshalEmptyString(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString("$0\r\n\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "" {
		t.Errorf("Expected empty string, got %v", result)
	}
}

func TestUnmarshalSimpleString(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString("+foobar\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	if result.(string) != "foobar" {
		t.Errorf("Expected foobar, got %v", result)
	}
}

func TestUnmarshalArray(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	array := result.([]Any)
	if len(array) != 2 {
		t.Errorf("Expected array of length 2, got %v", array)
	}
}

func TestUnmarshalEmptyArray(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString("*0\r\n")
	var result Any
	result, error := Unmarshal(buffer)
	if error != nil {
		t.Error(error)
	}
	array := result.([]Any)
	if len(array) != 0 {
		t.Errorf("Expected array of length 0, got %v", array)
	}
}
