package basic

import "fmt"

func slices() {
	//切片是一个可以动态变化数组，属于引用类型
	var array [5]int = [...]int{1, 2, 3, 4, 5}
	//定义一个切片
	slice := array[2:5] //引用数组下标从2到4的值(顾头不顾尾)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println("切片容量==>", cap(slice)) //切片容量是动态变化的

	//使用make创建切片(参数为 type,len,cap)
	ints := make([]int, 5, 6)
	ints[1] = 200
	fmt.Println(ints)

	//定义一个切片，直接制定具体数组
	strings := []string{"A", "B", "C", "D"}
	fmt.Println(strings)
	fmt.Println(len(strings))
	fmt.Println("切片容量==>", cap(strings))

	//动态追加具体元素
	strings01 := append(strings, "E", "F")
	fmt.Println(strings01)
	//也可以追加切片
	strings02 := append(strings01, strings...)
	fmt.Println(strings02)

	//切片拷贝
	strings03 := make([]string, 10)
	copy(strings03, strings01)
	fmt.Println(strings03)
	fmt.Println(len(strings03))
	fmt.Println("切片容量==>", cap(strings03))

	//string可以切片
	str := "helloWord!"
	slice01 := str[5:]
	fmt.Println(slice01)

	//进行部分替换
	bytes := []byte(str)
	bytes[0] = 'z'
	str = string(bytes)
	fmt.Println(str)

	//替换中文
	runes := []rune(str)
	runes[0] = '好'
	fmt.Println(string(runes))
}
