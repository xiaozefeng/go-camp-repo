package main

import (
	"log"
	"net"
)

func main() {

	var l, err = net.ListenUDP("udp", &net.UDPAddr{Port: 32000})
	if err != nil {
		log.Fatal(err)

	}

	defer l.Close()
	for {
		var buf [1024]byte
		n, addr, err := l.ReadFromUDP(buf[:])
		if err != nil {
			log.Printf("err = %+v\n", err)
			continue
		}
		data := append([]byte("hello \n"), buf[:n]...)
		l.WriteToUDP(data, addr)
	}

}
