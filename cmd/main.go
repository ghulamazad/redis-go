package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/ghulamazad/redis-clone/resp"
)

func main() {
	// Start a new  server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port :6379")

	// Listen from connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // close connection once finished

	for {
		deserializer := resp.NewDeserializer(conn)
		value, err := deserializer.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.Type != "array" {
			fmt.Println("Invalid request, exepected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, exepected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)
		handler, ok := resp.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Type: "string", Str: ""})
			continue
		}
		result := handler(args)
		writer.Write(result)
	}
}
