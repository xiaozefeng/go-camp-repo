package main

import (
	"bufio"
	"encoding/binary"
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

	// read
	protocol, err := decode(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("protocol = %+v\n", protocol)
}

func readPacketLen(conn net.Conn) (uint32, error) {
	return read(conn, 4)
}

func read(conn net.Conn, len int) (uint32, error) {
	var b = make([]byte, len)
	_, err := conn.Read(b)
	if err != nil {
		return 0, err
	}
	result := binary.BigEndian.Uint32(b)
	return result, nil
}

type Protocol struct {
	PacketLen uint32
	HeaderLen uint16
	Version   uint16
	Operation uint32
	Sequence  uint32
	Body      []byte
}

func (p *Protocol) ReadHeader(data []byte) {
	p.PacketLen = binary.BigEndian.Uint32(data[:4])
	p.HeaderLen = binary.BigEndian.Uint16(data[4:6])
	p.Version = binary.BigEndian.Uint16(data[6:8])
	p.Operation = binary.BigEndian.Uint32(data[8:12])
	p.Sequence = binary.BigEndian.Uint32(data[12:16])
}

func (p *Protocol) ReadBody(data []byte) {
	p.Body = data
}

func encode(body string) []byte {
	headerLen := 16
	packageLen := len(body) + headerLen

	ret := make([]byte, packageLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packageLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))

	version := 5
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation := 6
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))
	sequence := 7
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)
	return ret
}

func decode(conn net.Conn) (*Protocol, error) {
	data, err := readConn(conn, 16)
	if err != nil {
		return nil, err
	}

	var protocol = &Protocol{}
	protocol.ReadHeader(data)

	var bodyLen = int(protocol.PacketLen - 16)
	data, err = readConn(conn, bodyLen)
	if err != nil {
		return nil, err
	}
	protocol.ReadBody(data)

	return protocol, nil
}

func readConn(conn net.Conn, len int) ([]byte, error) {
	buf := make([]byte, len)
	_, err := bufio.NewReader(conn).Read(buf)
	return buf, err
}
