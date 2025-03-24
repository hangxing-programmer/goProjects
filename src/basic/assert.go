package basic

import "fmt"

type Point struct {
	x int
	y int
}

func main02() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //空接口，接受任何类型
	fmt.Println(a)

	var b Point
	//类型断言
	b = a.(Point)
	fmt.Println(b)
}
