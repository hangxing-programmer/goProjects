package main

import (
	"log"
	"net"
)

func main() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer udp.Close()

	buf := []byte("hello,udp")
	_, err = udp.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}

	buf = make([]byte, 1024)
	n, addr, err := udp.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("客户端接收到的数据:", string(buf[:n]), addr)

}
