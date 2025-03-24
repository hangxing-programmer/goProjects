package basic

import (
	"fmt"
	"time"
)

func startGCTask(stop chan bool) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	for {
		select {
		case <-t.C:
			fmt.Println("循环业务...", time.Now().Format("2006-01-02 15:04:05"))
		case <-stop:
			break
		}
	}
}
func main17() {
	stop := make(chan bool)
	go startGCTask(stop)
	time.Sleep(10 * time.Second)
	stop <- true
	close(stop)
}
