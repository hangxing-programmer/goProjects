package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	st = st.Elem()         //st指向的类型
	elem = reflect.New(st) //返回一个Value类型值，该值持有一个指针
	model = elem.Interface().(*user)
	elem = elem.Elem()                             //取得elem指向的值
	elem.FieldByName("UserId").SetString("123456") //赋值
	elem.FieldByName("Name").SetString("zs")
	fmt.Println("赋值后==>UserId", model.UserId)
	fmt.Println("赋值后==>Name", model.Name)
}

func main16() {

}
