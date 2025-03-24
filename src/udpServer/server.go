package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer udp.Close()
	for {
		var buf [1024]byte
		n, addr, err := udp.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("服务端接收到的数据", string(buf[:n]))
		//回复
		_, err = udp.WriteToUDP(buf[:n], addr)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
