package basic

import (
	"fmt"
	"sort"
)

func bubble(arr *[5]int) {
	fmt.Println("排序前==>", *arr)
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后==>", *arr)
}

// 排序分为内部和外部排序两种
func sorts() {

	//内部排序 —— 交换式排序 —— 01:冒泡排序:依次把最大排序冒泡至最上面
	arr := [5]int{100, 23, 56, 78, 90}
	bubble(&arr)

	var intSlice = []int{6, 4, 3, 8, 1}
	//1.冒泡排序
	//2.系统方法排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)

}
