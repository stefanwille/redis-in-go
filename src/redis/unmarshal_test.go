package redis

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func Unmarshal(buffer *bytes.Buffer) (result Any, error error) {
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
		lengthString, error := readLine(buffer)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		length, error := strconv.Atoi(lengthString)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		bytes := buffer.Next(length)
		s := string(bytes)
		_ = buffer.Next(2)
		return s, nil
	case '+':
		s, error := readLine(buffer)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		return s, nil
	case '*':
		lengthString, error := readLine(buffer)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		length, error := strconv.Atoi(lengthString)
		if error != nil {
			log.Fatal(error)
			return nil, error
		}
		var array []Any
		for i := 0; i < length; i++ {
			element, error := Unmarshal(buffer)
			if error != nil {
				log.Fatal(error)
				return nil, error
			}
			array = append(array, element)
		}
		return array, nil
	default:
		log.Fatalf("Unknown Redis type '%s'", string(r))
		return nil, fmt.Errorf("Unknown Redis type |'%s'|", string(r))
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
