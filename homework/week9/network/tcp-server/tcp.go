package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8110")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("err = %+v\n", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Printf("err = %+v\n", err)
			return
		}
		writer.WriteString("hello ")
		writer.Write(line)
		writer.Flush()
	}
}
