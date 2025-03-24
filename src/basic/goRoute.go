package basic

import (
	"fmt"
	"strconv"
	"time"
)

func testRoute() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello，goRoute" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main06() {

	go testRoute() //开启协程
	for i := 1; i <= 10; i++ {
		fmt.Println("hello，word" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
