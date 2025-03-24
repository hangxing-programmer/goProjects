package basic

import "fmt"

func WriteChannel(c1 chan int) {
	for i := 1; i <= 50; i++ {
		c1 <- i
		fmt.Println("写入数据...", i)
	}
	//关闭channel,后续可以for循环读取
	close(c1)
}
func ReadChannel(c1 chan int, c2 chan int) {
	for i := range c1 {
		c2 <- i
		fmt.Println("读取数据", i)
	}
	close(c2)
}

func main08() {

	channel01 := make(chan int, 50)
	channel02 := make(chan int, 50)
	go WriteChannel(channel01)
	go ReadChannel(channel01, channel02)

	//保证协程执行完
	for i := range channel02 {
		fmt.Println("读取数据===>", i)
	}

}
