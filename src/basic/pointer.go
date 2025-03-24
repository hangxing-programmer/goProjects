package basic

import "fmt"

func Pointer() {
	//指针类型
	var i int = 10
	fmt.Println("i的地址=", &i)

	//指针变量必须指向地址
	var ptr *int = &i
	fmt.Println("ptr=", ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

	var num int = 9
	var prt1 *int = &num
	*prt1 = 10
	fmt.Println("num=", num)

}
