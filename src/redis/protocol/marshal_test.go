package protocol

import (
	"bytes"
	"errors"
	"testing"
)

func marshal(any Any) string {
	var buffer bytes.Buffer

	error := Marshal(&buffer, any)
	if error != nil {
		panic(error)
	}
	return buffer.String()
}

func testMarshal(any Any, expected string, t *testing.T) {
	result := marshal(any)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMarshalInt(t *testing.T) {
	testMarshal(1000, ":1000\r\n", t)
}

func TestMarshalString(t *testing.T) {
	testMarshal("foobar", "$6\r\nfoobar\r\n", t)
}

func TestMarshalEmptyString(t *testing.T) {
	testMarshal("", "$0\r\n\r\n", t)
}

func TestMarshalError(t *testing.T) {
	testMarshal(errors.New("Error message"), "-Error message\r\n", t)
}

func TestMarshalNil(t *testing.T) {
	testMarshal(nil, "$-1\r\n", t)
}

func TestMarshalArray(t *testing.T) {
	var a []Any = []Any{"foo", "bar"}
	testMarshal(&a, "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n", t)
}

func TestMarshalEmptyArray(t *testing.T) {
	var a []Any = []Any{}
	testMarshal(&a, "*0\r\n", t)
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

func TestMarshalSimpleString(t *testing.T) {
	var b bytes.Buffer

	MarshalSimpleString(&b, "foobar")
	var result string = b.String()
	expected := "+foobar\r\n"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
