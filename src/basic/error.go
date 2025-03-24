package basic

import (
	"fmt"
	"time"
)

func Write() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("输出数据...", i)
	}
}
func Read() {
	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误...", err)
		}
	}()
	var m map[int]string
	//没有make，会报错
	m[0] = "zs"
}

func main05() {

	go Write()
	go Read()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("ok")
	}
}
