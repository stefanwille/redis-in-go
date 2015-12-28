package protocol

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func Unmarshal(reader *bufio.Reader) (result Any, error error) {
	r, _, error := reader.ReadRune()
	if error != nil {
		return
	}
	switch r {
	case ':':
		integerString, error := readLine(reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		i, error := strconv.Atoi(integerString)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		return i, nil
	case '$':
		lengthString, error := readLine(reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		length, error := strconv.Atoi(lengthString)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		bytes, error := next(length, reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		_, error = next(2, reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		s := string(*bytes)
		return s, nil
	case '+':
		s, error := readLine(reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		return s, nil
	case '*':
		lengthString, error := readLine(reader)
		if error != nil {
			log.Print(error)
			return nil, error
		}
		length, error := strconv.Atoi(lengthString)
		if error != nil {
			return nil, error
		}
		var array []Any
		for i := 0; i < length; i++ {
			element, error := Unmarshal(reader)
			if error != nil {
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

// Reads the next n bytes
func next(n int, reader *bufio.Reader) (bytes *[]byte, error error) {
	buffer := make([]byte, n)
	nread, error := reader.Read(buffer)
	if error != nil {
		return nil, error
	}
	if nread != n {
		return nil, fmt.Errorf("next(%d) received only %d bytes", n, nread)
	}
	return &buffer, nil
}

func readLine(buffer *bufio.Reader) (line string, error error) {
	line = ""
	for {
		c, _, error := buffer.ReadRune()
		if error != nil {
			log.Print(error)
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
		log.Print(error)
		return line, error
	}
	if c != '\n' {
		log.Print("Expected newline")
		return line, fmt.Errorf("Expected newline, got %s", string(c))
	}
	return line, nil
}
