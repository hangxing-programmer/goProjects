package basic

import (
	"fmt"
	"log"
	"os"
)

func create(target, linkName string) {
	// 创建目标文件（如果不存在）
	if _, err := os.Stat(target); os.IsNotExist(err) {
		if err := os.WriteFile(target, []byte("test content"), 0644); err != nil {
			log.Fatal(err)
		}
	}
	// 检查符号链接是否存在，如果存在则删除
	if stat(linkName) {
		remove(linkName)
	}
	// 创建符号链接
	if err := os.Symlink(target, linkName); err != nil {
		log.Fatal("创建符号链接失败:", err)
	}

	fmt.Println("符号链接创建成功:", linkName, "->", target)

	resolved, err := os.Readlink(linkName)
	if err != nil {
		log.Fatal("读取符号链接失败:", err)
	}
	fmt.Println("符号链接指向:", resolved) // 应该输出 "test_file.txt"
}
func read(linkName string) {
	// 通过符号链接读取文件内容
	content, err := os.ReadFile(linkName)
	if err != nil {
		log.Fatal("读取符号链接文件失败:", err)
	}
	fmt.Println("符号链接文件内容:", string(content))
}
func write(linkName string, content []byte) {
	file, err := os.OpenFile(linkName, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("打开符号链接文件失败:", err)
	}
	defer file.Close()

	if _, err := file.Write(content); err != nil {
		log.Fatal("通过符号链接写入文件失败:", err)
	}

	fmt.Println("通过符号链接写入文件成功")
}
func remove(linkName string) {
	// 删除符号链接
	if err := os.Remove(linkName); err != nil {
		log.Fatal("删除符号链接失败:", err)
	}

	fmt.Println("符号链接删除成功")
}
func stat(linkName string) bool {
	// 检查符号链接是否存在
	if _, err := os.Lstat(linkName); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func main() {
	//target := "C:\\temp\\test_file.txt" // 目标文件
	linkName := "testlink" // 符号链接名
	//create(target, linkName)
	//write(linkName, []byte("testing"))
	read(linkName)

}
