package basic

import (
	"fmt"
	"time"
)

func getTime() {
	now := time.Now()
	fmt.Println("当前时间:", now)
	//获取年月日时分秒
	fmt.Println("年", now.Year())
	fmt.Println("月", int(now.Month()))
	fmt.Println("日", now.Day())
	fmt.Println("时", now.Hour())
	fmt.Println("分", now.Minute())
	fmt.Println("秒", now.Second())
	//日期格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	//时间常量(time.)
	fmt.Println("1微秒=", time.Microsecond)
	//休眠
	fmt.Println("休眠开始...")
	time.Sleep(time.Second * 1)
	fmt.Println("休眠结束...")
	//获取当前时间戳
	fmt.Printf("当前时间戳===>%v,纳秒级===>%v\n", now.Unix(), now.UnixNano())
}
