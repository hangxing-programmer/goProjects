package basic

//
//import (
//	"fmt"
//	"io"
//	"sync"
//)
//
//// PipeReader 结构体，包含一个 *io.PipeReader 和一个等待函数
//type PipeReader struct {
//	*io.PipeReader
//	wait func()
//}
//
//// PipeWriter 结构体，包含一个 *io.PipeWriter 和一个标记完成的函数
//type PipeWriter struct {
//	*io.PipeWriter
//	done func()
//}
//
//// WaitPipe 函数创建一个管道，并返回 PipeReader 和 PipeWriter 的实例
//func WaitPipe() (*PipeReader, *PipeWriter) {
//	r, w := io.Pipe()
//	var wg sync.WaitGroup
//	wg.Add(1)
//	return &PipeReader{
//			PipeReader: r,
//			wait:       wg.Wait,
//		}, &PipeWriter{
//			PipeWriter: w,
//			done:       wg.Done,
//		}
//}
//
//func main() {
//	// 创建 PipeReader 和 PipeWriter
//	reader, writer := WaitPipe()
//	go func() {
//		defer writer.Close()
//		n, err := writer.Write([]byte("hello world"))
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("wrote %d bytes\n", n)
//		writer.done()
//	}()
//
//	bytes, err := io.ReadAll(reader)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("read %d bytes, data = %s\n", len(bytes), string(bytes))
//	reader.wait()
//	fmt.Println("over...")
//}
