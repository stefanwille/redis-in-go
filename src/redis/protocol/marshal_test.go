package protocol

import (
	"bytes"
	"errors"
	"testing"
)

func TestMarshalInt(t *testing.T) {
	var buffer bytes.Buffer

	Marshal(&buffer, 1000)
	var result string = buffer.String()
	if result != ":1000\r\n" {
		t.Error(result)
	}
}

func TestMarshalString(t *testing.T) {
	var b bytes.Buffer

	Marshal(&b, "foobar")
	var result string = b.String()
	if result != "$6\r\nfoobar\r\n" {
		t.Error(result)
	}
}

func TestMarshalEmptyString(t *testing.T) {
	var b bytes.Buffer

	Marshal(&b, "")
	var result string = b.String()
	if result != "$0\r\n\r\n" {
		t.Error(result)
	}
}

func TestMarshalError(t *testing.T) {
	var b bytes.Buffer

	value := errors.New("Error message")
	error := Marshal(&b, value)
	if error != nil {
		t.Errorf("Got error: %v", error)
		return
	}
	var result string = b.String()
	if result != "-Error message\r\n" {
		t.Error(result)
	}
}

func TestMarshalNil(t *testing.T) {
	var b bytes.Buffer

	error := Marshal(&b, nil)
	if error != nil {
		t.Errorf("Got error: %v", error)
		return
	}
	var result string = b.String()
	if result != "$-1\r\n" {
		t.Error(result)
	}
}

func TestMarshalArray(t *testing.T) {
	var b bytes.Buffer

	var a []Any = []Any{"foo", "bar"}

	Marshal(&b, &a)
	var result string = b.String()
	expected := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMarshalEmptyArray(t *testing.T) {
	var b bytes.Buffer

	var a []Any = []Any{}

	Marshal(&b, &a)
	var result string = b.String()
	expected := "*0\r\n"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMarshalSimpleString(t *testing.T) {
	var b bytes.Buffer

	MarshalSimpleString(&b, "foobar")
	var result string = b.String()
	expected := "+foobar\r\n"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
func TestMarshalUnsupportedTypeReturnsAnError(t *testing.T) {
	var b bytes.Buffer

	var m map[string]string
	var a []Any = []Any{m}

	error := Marshal(&b, &a)
	if error == nil {
		t.Error("Expected error, got nil")
	}
}
