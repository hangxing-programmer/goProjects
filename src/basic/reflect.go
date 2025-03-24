package basic

import (
	"fmt"
	"reflect"
)

func getReflect(b interface{}) {
	//1.获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)
	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	kind := rVal.Kind()
	fmt.Println("rVal的类别==>", kind, ";rVal的值==>", rVal)
	//虽然传值是int,但是不能和int类型进行加减等操作
	fmt.Printf("%T %v\n", rType, rType)
	//如果进行加减等操作，需要转型
	num := 2 + rVal.Int()
	fmt.Println("可以进行加法操作==>", num)
	//下面将rVal转换成interface{}
	iV := rVal.Interface()
	//再进行断言
	num1 := iV.(int)
	fmt.Println("转换完成==>", num1)
}

func main13() {
	num := 100
	getReflect(num)
}
