package basic

import (
	"flag"
	"fmt"
	"os"
)

func getArgs() {

	//1.有序获取命令行参数
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	//2.无序获取命令行参数
	var user string
	var pwd string
	var host string
	var port int

	//&user接收参数 -u 后的值
	//"u" -u指定参数
	//" " 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "地址,默认为空")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	flag.Parse()

	fmt.Printf("user: %s, pwd: %s, host: %s, port: %d\n", user, pwd, host, port)
}
