package main

import (
	"fmt"
	"net"

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
		_ = value

		writer := resp.NewWriter(conn)
		writer.Write(resp.Value{Type: "string", Str: "OK"})
	}
}
