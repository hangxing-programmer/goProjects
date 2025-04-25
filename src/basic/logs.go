package basic

import (
	"io"
	"log"
	"os"
)

func main() {
	// 打开或创建日志文件 a.txt
	fileA, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open a.txt: %v", err)
	}
	defer fileA.Close()

	// 打开或创建日志文件 b.txt
	fileB, err := os.OpenFile("test02.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open b.txt: %v", err)
	}
	defer fileB.Close()

	// 创建一个多路复用器，将日志输出到两个文件
	multiWriter := io.MultiWriter(fileA, fileB)

	// 创建一个日志记录器，输出到多路复用器
	log.SetOutput(multiWriter)

	// 输出日志到两个文件
	log.Println("This log message will be written to both a.txt and b.txt")

	// 创建两个独立的日志记录器，分别输出到 a.txt 和 b.txt
	loggerA := log.New(fileA, "LoggerA: ", log.LstdFlags|log.Lshortfile)
	loggerB := log.New(fileB, "LoggerB: ", log.LstdFlags|log.Lshortfile)

	// 使用独立的日志记录器输出日志
	loggerA.Println("This log message will only be written to a.txt")
	loggerB.Println("This log message will only be written to b.txt")
}
