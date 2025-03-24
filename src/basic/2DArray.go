package basic

import "fmt"

func array() {
	//定义二维数组
	var arr [4][6]int
	arr[1][2] = 1
	arr[3][0] = 1
	arr[0][2] = 1
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	arr01 := [2][3]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(arr01)
}
