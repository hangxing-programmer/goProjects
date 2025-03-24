package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送的消息
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送消息,如果客户端没有Write,该协程会阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务端Read Err===>", err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
	//及时关闭连接
	defer conn.Close()
}
func main() {
	fmt.Println("服务器开始监听...")
	//使用网络协议tcp,在本地监听8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败===>", err)
		return
	}
	defer listen.Close() //及时关闭资源
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败===>", err)
			return
		} else {
			fmt.Println(conn.RemoteAddr().String(), "连接成功...", conn)
		}
		//为客户端服务
		go process(conn)
	}
}
