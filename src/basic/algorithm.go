package basic

import (
	"bufio"
	"os"
	"strconv"
)

type Val struct {
	row   int
	col   int
	value int
}

func main03() {
	//原始数组
	var m [11][11]int
	m[2][2] = 1
	m[3][3] = 2

	val := Val{
		row:   11,
		col:   11,
		value: 0,
	}
	//稀疏数组
	var sparseArry []Val
	sparseArry = append(sparseArry, val)
	for i, v := range m {
		for i2, v2 := range v {
			if v2 != 0 {
				sparseArry = append(sparseArry, Val{
					row:   i,
					col:   i2,
					value: v2,
				})
			}
		}
	}

	open, _ := os.OpenFile("D:\\test\\chess.data", os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(open)
	defer open.Close()
	for i, v := range sparseArry {
		if i != 0 {
			writer.WriteString(strconv.Itoa(v.row) + "\t" + strconv.Itoa(v.col) + "\t" + strconv.Itoa(v.value) + "\n")
		}
	}
	writer.Flush()
}
