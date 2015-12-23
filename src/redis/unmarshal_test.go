package redis

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func Unmarshal(buffer *bytes.Buffer, index int) (result Any, error error) {
	r, _, error := buffer.ReadRune()
	if error != nil {
		log.Fatal(error)
	}
	switch r {
	case ':':
		integerString, error := readLine(buffer)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		i, error := strconv.Atoi(integerString)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		return i, nil
	case '$':
		return nil, nil
	default:
		log.Fatalf("Unknown Redis type %s", string(r))
		return nil, fmt.Errorf("Unknown Redis type")
	}
}

func readLine(buffer *bytes.Buffer) (line string, error error) {
	line = ""
	for {
		c, _, error := buffer.ReadRune()
		if error != nil {
			log.Fatal(error)
			return line, error
		}

		if c == '\r' {
			break
		} else {
			line += string(c)
		}
	}

	// Read newline
	c, _, error := buffer.ReadRune()
	if error != nil {
		log.Fatal(error)
		return line, error
	}
	if c != '\n' {
		log.Fatal("Expected newline")
		return line, fmt.Errorf("Expected newline, got %s", string(c))
	}
	return line, nil
}

func TestUnmarshalInt(t *testing.T) {
	var buffer *bytes.Buffer = bytes.NewBufferString(":1000\r\n")
	var result Any
	result, error := Unmarshal(buffer, 0)
	if error != nil {
		t.Error(error)
	}
	if result.(int) != 1000 {
		t.Errorf("Expected 1000, got %v", result)
	}
}

// func TestUnmarshalString(t *testing.T) {
// 	var buffer *bytes.Buffer = bytes.NewBufferString("$6\r\nfoobar\r\n")
// 	var result Any
// 	result, _, error := Unmarshal(buffer, 0)
// 	if error != nil {
// 		t.Error(error)
// 	}
// 	if result.(string) != "foobar" {
// 		t.Errorf("Expected foobar, got %v", result)
// 	}
// }

// func TestMarshalString(t *testing.T) {
// 	var b bytes.Buffer

// 	Marshal(&b, "foobar")
// 	var result string = b.String()
// 	if result != "$6\r\nfoobar\r\n" {
// 		t.Error(result)
// 	}
// }

// func TestMarshalEmptyString(t *testing.T) {
// 	var b bytes.Buffer

// 	Marshal(&b, "")
// 	var result string = b.String()
// 	if result != "$0\r\n\r\n" {
// 		t.Error(result)
// 	}
// }

// func TestMarshalArray(t *testing.T) {
// 	var b bytes.Buffer

// 	var a []Any = []Any{"foo", "bar"}

// 	Marshal(&b, &a)
// 	var result string = b.String()
// 	expected := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestMarshalEmptyArray(t *testing.T) {
// 	var b bytes.Buffer

// 	var a []Any = []Any{}

// 	Marshal(&b, &a)
// 	var result string = b.String()
// 	expected := "*0\r\n"
// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestMarshalSimpleString(t *testing.T) {
// 	var b bytes.Buffer

// 	MarshalSimpleString(&b, "foobar")
// 	var result string = b.String()
// 	expected := "+foobar\r\n"
// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }
