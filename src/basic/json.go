package basic

import (
	"encoding/json"
	"fmt"
)

type Person01 struct {
	Name   string
	Age    int
	Salary float64
}

// 结构体序列化
func testStruct() {
	person := Person01{
		Name:   "小白",
		Age:    18,
		Salary: 9999.99,
	}
	//JSON序列化
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

// 数组序列化
func testMap() {
	a := make(map[string]string)
	a["name"] = "小白"
	a["age"] = "18"
	a["salary"] = "9999.99"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

// 切片序列化
func testSlice() {
	a := make(map[string]string)
	a["name"] = "小白"
	a["age"] = "18"
	a["salary"] = "9999.99"
	var slice []map[string]string
	slice = append(slice, a)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

type Person02 struct {
	Name   string
	Age    int
	Salary float64
}

// 反序列化成struct
func ujson() {
	str := "{\"age\":18,\"name\":\"小白\",\"salary\":9999.99}"
	var person Person02
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)
}

func main11() {
	//基本数据类型序列化,没有实际意义
	data, err := json.Marshal(234.56)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))

	//decoder := json.NewDecoder(os.Stdin)
	//encoder := json.NewEncoder(os.Stdout)
	//for {
	//	var v map[string]interface{}
	//	if err := decoder.Decode(&v); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	for i := range v {
	//		if i != "title" {
	//			delete(v, i)
	//		}
	//	}
	//	if err := encoder.Encode(v); err != nil {
	//		log.Println(err)
	//	}
	//}
}
