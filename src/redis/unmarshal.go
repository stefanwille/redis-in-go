package redis

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
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
