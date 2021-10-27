package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := "hello"
	encoded := encode(input)
	// send
	fmt.Fprintln(conn, encoded)

	line, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	decoded := decode(line)

}

func encode(body string) []byte {
	return nil
}

func decode(data []byte) string {
	return ""
}
