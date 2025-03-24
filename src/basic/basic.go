package basic

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

// 定义全局变量
var (
	gloabl = 0
)

func basic() {
	//c, d := utils.Add(1, 2, "hello")
	//basic.Pointer()
	//fmt.Printf("c=%d , d=%s\n", c, d)

	//匿名函数(也可以放到全局变量，函数名大写实现全局通用,匿名函数后必须跟着小括号!!!)
	//sum := func(n1 int, n2 int) int {
	//	return n1 + n2
	//}
	//result := sum(100, 200)
	//fmt.Printf("sum=%d\n", result)
	//内置函数
	num1 := 100
	fmt.Printf("类型:%t,值:%v,地址:%v\n", num1, num1, &num1)
	//new分配值类型内存
	num2 := new(int)
	fmt.Printf("类型:%t,值:%v,地址:%v\n", num2, num2, &num2)
	//随机数
	//n := rand.Intn(10) + 1
	//fmt.Println(n)
	//var str string = "qwer!北京"
	//for index, val := range str {
	//	fmt.Printf("index=%d,val=%c \n", index, val)
	//}

	//label1:
	//	for i := 0; i <= 10; i++ {
	//		fmt.Println("今天星期四")
	//		for j := 0; j <= 5; j++ {
	//			if j == 2 {
	//				fmt.Println("=============》", j)
	//				break label1
	//			}
	//		}
	//	}

	//var a int
	//fmt.Print("请输入一个字符:")
	//fmt.Scanf("%d", &a)
	//switch a {
	//case 1:
	//	fmt.Println("星期一")
	//	fallthrough //默认穿透一层
	//case 2:
	//	fmt.Println("星期二")
	//case 3:
	//	fmt.Println("星期三")
	//default:
	//	fmt.Println("周日")
	//}

	//-2的原码 1000 0010 反码 1111 1101 补码 1111 1110
	//2的补码 0000 0010
	//补码比对  1111 1110
	//补码比对  0000 0010
	fmt.Println(-2 & 2)
	fmt.Println(-2 | 2)
	fmt.Println(-2 ^ 2) //1111 1100 =>补码 1111 1011 =>原码 1000 0100 = -4

	//右移运算
	m := 1 >> 2 //0000 0001 => 0000 0000 = 0
	//左移运算
	n := 1 << 2 //0000 0001 => 0000 0100 = 4
	fmt.Println(m, n)
	//二进制
	var number int = 5
	fmt.Printf("%b\n", number)
	//八进制(以数字0开头)
	var number1 int = 011
	fmt.Println(number1)
	//十六进制
	var number2 int = 0x11
	fmt.Println(number2)

	var name string
	fmt.Printf("请输入名字：")
	fmt.Scanln(&name)
	fmt.Println(name)

	var s string = "true"
	var i bool
	i, _ = strconv.ParseBool(s)
	fmt.Printf("转换类型后%v\n", i)
	//整形转string类型
	var str = fmt.Sprintf("%d", 1)
	fmt.Printf(str + "\n")
	//布尔默认值
	var h bool
	fmt.Println(h)
	var g = "ddddddddddddd"
	fmt.Println(g)
	var f byte = 'a'
	fmt.Printf("f=%c\n", f)

	//科学计数法
	var e = 1.2323e3
	fmt.Println(e)

	var c float32 = 1.2344523101001010101111111111111
	var d = 1.2344523101001010101111111111111
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("global\t", gloabl)
	var a byte = 255
	fmt.Printf("a的类型是%T\n", a)
	var b int64 = 10
	fmt.Printf("b占用的字节数是%d\n", unsafe.Sizeof(b))

	//fmt.Println("helloWord")
	//var i = 2
	//num := 3
	//i = 4
	//fmt.Println(i)
	//fmt.Println("num\n", num)
	//
	//var a, b, c int
	//fmt.Println(a, b, c)
	//var abc = "adc"
	//fmt.Println(abc)

	//获取CPU数
	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)

	//定义常量
	const (
		a1     = iota       //常量0
		b1                  //0+1
		c1, d1 = iota, iota //0+1+1
	)
	fmt.Println(a1, b1, c1)
}
