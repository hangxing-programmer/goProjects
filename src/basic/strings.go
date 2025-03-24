package basic

import (
	"fmt"
	"strconv"
	"strings"
)

func str() {
	var str = "hello北京"
	//遍历字符串
	str2 := []rune(str)
	fmt.Printf("%c\n", str2[6])
	//字符串转整数
	n, err := strconv.Atoi("123")
	fmt.Printf("n=%v,err是%v\n", n, err)
	//整数转字符串
	str = strconv.Itoa(n)
	fmt.Println(str)
	//字符串转[]byte
	var bytes = []byte(str)
	//[]byte转字符串
	str = string(bytes)
	fmt.Println("转换后的字符串===>", str)
	//10进制转2,8,16,返回对应的字符串
	fmt.Println("123对应的二进制", strconv.FormatInt(123, 2))
	fmt.Println("123对应的八进制", strconv.FormatInt(123, 8))
	fmt.Println("123对应的十六进制", strconv.FormatInt(123, 16))
	//验证是否包含指定字符串
	contains := strings.Contains("he", str)
	fmt.Println(contains)

	str = "hehehello"
	//有几个指定字符串
	count := strings.Count(str, "he")
	fmt.Println(count)
	//不区分大小写(==区分大小写)
	fold := strings.EqualFold("hello", "Hello")
	fmt.Println(fold)
	//第一次出现指定字符串的索引位置,如果没有返回-1
	index := strings.Index(str, "he")
	fmt.Println(index)
	//最后一次出现
	lastIndex := strings.LastIndex(str, "he")
	fmt.Println(lastIndex)
	//替换指定字符串(1代表替换几个,n代表全部)
	replace := strings.Replace(str, "h", "H", 1)
	fmt.Println(replace)
	//分割字符串为数组(指定分隔标识)
	split := strings.Split(str, "")
	fmt.Println(split)
	//大小写转换
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	//去掉字符串左右两边的空格
	str = " hero hero hero "
	fmt.Println(strings.TrimSpace(str))
	//将字符串左右两边指定字符串去掉
	str = " hero!! hero!! hero!! "
	fmt.Println(strings.Trim(str, " !"))
	//只去除左边
	str = " !!hero!! hero!! hero!! "
	fmt.Println(strings.TrimLeft(str, " !"))
	//只去除右边
	str = " hero!! hero!! hero!! "
	fmt.Println(strings.TrimRight(str, " !"))
	//是否以指定字符串开头,结尾
	fmt.Println(strings.HasPrefix(str, " hero"))
	fmt.Println(strings.HasSuffix(str, "hero"))
}
