package basic

import (
	"fmt"
	"math/rand"
)

func arr() {
	//数组是值类型
	//四种初始化数组方式
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println(array)
	fmt.Printf("数组地址===>%p\n", &array)

	var array01 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array01)

	var array02 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(array02)

	var array03 = [5]int{1: 100, 2: 200, 3: 300, 4: 400}
	fmt.Println(array03)

	//数组遍历(如果不要index，可以换为_)
	for index, value := range array {
		fmt.Println(index, value)
	}

	var chars [26]byte
	for i := 0; i < 26; i++ {
		chars[i] = 'A' + byte(i)
		fmt.Printf("%c ", chars[i])
	}

	//数组反转
	var array04 [5]int
	for i := 0; i < len(array04); i++ {
		array04[i] = rand.Intn(100)
	}
	fmt.Println("交换前:", array04)
	//临时变量
	temp := 0
	for i := 0; i < len(array04)/2; i++ {
		temp = array04[len(array04)-1-i]
		array04[len(array04)-1-i] = array04[i]
		array04[i] = temp
	}
	fmt.Println("交换后:", array04)
}
