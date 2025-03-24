package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("客户端连接失败===>", err)
		return
	}
	fmt.Println("连接成功...", conn)

	//功能1.客户端可已发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入[终端]
	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败===>", err)
		}
		//输入exit,退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Print("客户端退出...")
			break
		}
		//再将line 发送给服务器
		n, err := conn.Write([]byte(line + "\r\n"))
		if err != nil {
			fmt.Println("发送失败===>", err)
		}
		fmt.Println("客户端发送字节数:", n)
	}

}
