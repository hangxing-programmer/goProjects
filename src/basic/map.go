package basic

import "fmt"

func map01() {
	//定义map，必须要make，分配数据空间
	var map01 map[string]string
	map01 = make(map[string]string)
	map01["a"] = "A"
	map01["b"] = "B"
	map01["a"] = "C" //同一key的value会被覆盖
	fmt.Println(map01)

	map02 := make(map[string]string)
	map02["a"] = "A"
	map02["b"] = "B"
	map02["c"] = "C"
	//删除
	delete(map02, "a")
	//查找
	_, s := map02["d"]
	if s {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(map02)

	//遍历
	for k, v := range map02 {
		fmt.Println(k, v)
	}

	//切片
	map03 := make(map[string]string)
	map03["上"] = "北"
	map03["下"] = "南"
	map03["左"] = "西"
	map03["右"] = "东"
	//make切片
	mapSilce := make([]map[string]string, 3)
	//复制
	mapSilce[0] = make(map[string]string)
	mapSilce[0]["上"] = map03["上"]
	fmt.Println(mapSilce)
}
