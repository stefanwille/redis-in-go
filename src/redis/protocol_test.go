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
	default:
		panic(fmt.Sprintf("Unrecognized type '%T' for redis.Marshal", t))
	}
}

func TestMarshalInt(t *testing.T) {
	var b bytes.Buffer

	Marshal(&b, 1000)
	var result string = b.String()
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
