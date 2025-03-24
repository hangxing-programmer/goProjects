package basic

import (
	"fmt"
	"sync"
	"time"
)

var (
	m = make(map[int]int)
	//声明一个全局互斥锁
	// lock:全局互斥锁  sync:包,同步  Mutex:互斥
	lock sync.Mutex
)

func add(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	m[n] = res
	//解锁
	lock.Unlock()
}

func main07() {
	for i := 0; i < 10; i++ {
		go add(i)
	}

	time.Sleep(time.Second * 5)

	for i, v := range m {
		fmt.Println(i, v)
	}

}
