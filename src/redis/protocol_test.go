package redis

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type Any interface{}

func Marshal(writer io.Writer, value Any) {
	switch t := value.(type) {
	case int:
		var intValue int = value.(int)
		fmt.Fprintf(writer, ":%d\r\n", intValue)
	case string:
		var stringValue string = value.(string)
		fmt.Fprintf(writer, "$%d\r\n%s\r\n", len(stringValue), stringValue)
	case *[]Any:
		var arrayPointerValue *[]Any = value.(*[]Any)
		fmt.Fprintf(writer, "*%d\r\n", len(*arrayPointerValue))
		for _, anyValue := range *arrayPointerValue {
			Marshal(writer, anyValue)
		}
	case []Any:
		var arrayValue []Any = value.([]Any)
		Marshal(writer, &arrayValue)
	default:
		panic(fmt.Sprintf("Unrecognized type '%T' for redis.Marshal", t))
	}
}

func MarshalSimpleString(writer io.Writer, value string) {
	fmt.Fprintf(writer, "+%s\r\n", value)
}

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
