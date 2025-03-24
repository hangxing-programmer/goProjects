package basic

import (
	"fmt"
	"time"
)

type Cat01 struct {
	Name string
	Age  int
}

func main04() {

	//创建一个可以存放3个int类型的channel
	var intChan chan int
	intChan = make(chan int, 3)

	//写入数据
	intChan <- 12
	num := 211
	intChan <- num

	//channel的长度和容量
	fmt.Println(len(intChan), cap(intChan))

	//读取数据,先进先出!!!
	outNum := <-intChan
	fmt.Println(outNum)
	fmt.Println(len(intChan), cap(intChan))

	//遍历
	all := make(chan interface{}, 6)
	all <- 10
	all <- Cat01{"黑猫", 10}
	all <- 11
	all <- 12
	all <- 13
	all <- 14
	//关闭channel,不可以再次写入,可以读取
	close(all)
	//遍历channel,必须要先关闭channel!!!
	for v := range all {
		fmt.Println(v)
	}

	ready := make(chan int)
	go func() {
		fmt.Println("通道打开中...")
		time.Sleep(3 * time.Second)
		//关闭ready通道来通知主 goroutine。主 goroutine 从通道接收数据（实际上是等待通道关闭），然后继续执行后续的代码
		close(ready)
	}()
	<-ready
	fmt.Println("It's over")

}
