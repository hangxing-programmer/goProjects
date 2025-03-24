package basic

import (
	"fmt"
	"reflect"
)

type stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 反射调用，结构体所绑定的方法名必须首字母大写！！！
func (student stu) GetMsg(n1 int) {
	fmt.Println("反射调用成功!!!返回值==>", n1)
}

func getReflect00(b interface{}) {
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
func getReflect01(b interface{}) {
	//通过指针，进而修改值
	rVal := reflect.ValueOf(b)
	kind := rVal.Kind()
	fmt.Println("b的类别==>", kind)
	//.Elem()是对value的封装
	rVal.Elem().SetInt(20)
}
func getReflect02(b interface{}) {
	rVal := reflect.ValueOf(b)
	numField := rVal.NumField()
	fmt.Println("该结构体字段数:", numField)
}
func testStruct01(s interface{}) {
	//声明切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf(11))

	//调用方法
	rVal := reflect.ValueOf(s)
	rVal.Method(0).Call(params)

	//查询json参数
	rType := reflect.TypeOf(s)
	num := rVal.NumField()
	for i := 0; i < num; i++ {
		res := rType.Field(i).Tag.Get("json")
		fmt.Printf("第%d个参数是%v\n", i, res)
	}

}
func main15() {
	//num := 100
	//getReflect00(num)
	//getReflect01(&num)
	//fmt.Println("修改后的值:", num)

	stu01 := stu{"zs", 10}
	//getReflect02(stu01)

	fmt.Println("反射调用方法开始...")
	testStruct01(stu01)
	fmt.Println("反射调用方法结束...")
}
