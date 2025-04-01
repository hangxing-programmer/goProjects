package basic

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// 记录文件内数字，字母，空格，其他的个数
//
//	type nums struct {
//		nums    int
//		letters int
//		space   int
//		other   int
//	}
func listFiles() {
	root := "D:\\vol2"
	open, _ := os.Open(root)
	entries, _ := open.ReadDir(-1)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(root + "\\" + entry.Name())
			err := filepath.WalkDir(root+"\\"+entry.Name(), func(path string, dirEntry fs.DirEntry, err error) error {
				if err != nil {
					fmt.Printf("访问 %q 失败: %v\n", path, err)
					return nil // 跳过错误，继续遍历
				}
				if !dirEntry.IsDir() { // 只处理文件，不处理目录
					fmt.Println(path) // 直接输出完整路径
				}
				return nil
			})

			if err != nil {
				fmt.Printf("遍历目录失败: %v\n", err)
			}
		}
		fmt.Println(root + "\\" + entry.Name())

	}
}
func readFile() {

	filePath := "D:\\userTest\\test01.txt"
	filePath01 := "D:\\userTest\\test01.txt"
	_, err2 := os.Stat(filePath)
	if err2 == nil {
		fmt.Println("文件存在...")
	}
	//打开文件，向文件内添加数据
	file, err := os.OpenFile(filePath01, os.O_WRONLY|os.O_APPEND, 0666) //perm在windows不起作用
	if err == nil {
		writer := bufio.NewWriter(file)
		for i := 0; i < 5; i++ {
			writer.WriteString("hello,go!\n")
		}
		//因为writer是带缓存的，内容先写入到缓存
		writer.Flush()
	}

	//读取文件
	readFile, err := os.ReadFile(filePath)
	if err == nil {
		fmt.Println(string(readFile))
	}
	err = os.WriteFile(filePath01, readFile, 0666)
	if err == nil {
		fmt.Println(string(readFile))
	}

	//及时关闭file
	defer file.Close()

	//num := nums{}

	//file := "D:\\userTest\\userTest.txt"
	//readFile, _ := os.ReadFile(file)
	//for _, i := range readFile {
	//	switch {
	//	case i >= 'A' && i <= 'Z':
	//		fallthrough //穿透
	//	case i >= 'a' && i <= 'z':
	//		num.letters++
	//	case i >= '0' && i <= '9':
	//		num.nums++
	//	case i == ' ' || i == '\t' || i == '\n' || i == '\r':
	//		num.space++
	//	default:
	//		num.other++
	//	}
	//}
	//fmt.Println("字母总数=", num.letters, "数字总数=", num.nums, "空格总数=", num.space, "其他总数=", num.other)

	////拷贝图片
	//file := "D:\\pics\\pic.png"
	//target := "D:\\userTest"
	//
	//openFile, _ := os.OpenFile(file, os.O_RDONLY, 0666)
	//reader := bufio.NewReader(openFile)
	//
	//targetFile, _ := os.OpenFile(target, os.O_RDONLY, 0666)
	//writer := bufio.NewWriter(targetFile)
	//
	////执行拷贝
	//_, err := io.Copy(writer, reader)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("拷贝完成...")
	//}
	//
	//defer openFile.Close()
	//defer targetFile.Close()
}
