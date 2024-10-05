package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Deserializer struct {
	reader *bufio.Reader
}

func NewDeserializer(rd io.Reader) *Deserializer {
	return &Deserializer{reader: bufio.NewReader(rd)}
}

func (r *Deserializer) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *Deserializer) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Deserializer) Read() (Value, error) {
	_type, err := r.reader.ReadByte()

	if err != nil {
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(_type))
		return Value{}, nil
	}
}

func (r *Deserializer) readArray() (Value, error) {
	v := Value{}
	v.Type = "array"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	// foreach line, parse and read the value
	v.Array = make([]Value, 0)

	for i := 0; i < len; i++ {
		val, err := r.Read()

		if err != nil {
			return v, err
		}

		v.Array = append(v.Array, val)
	}
	return v, nil
}

func (r *Deserializer) readBulk() (Value, error) {
	v := Value{}
	v.Type = "bulk"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}
	bulk := make([]byte, len)

	r.reader.Read(bulk)
	v.Bulk = string(bulk)

	// Read the trailing CRLF
	r.readLine()

	return v, nil
}
