package basic

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// PipeReader 结构体，包含一个 *io.PipeReader 和一个等待函数
type PipeReader struct {
	*io.PipeReader
	wait     func()
	buffered *bufio.Reader
}

// PipeWriter 结构体，包含一个 *io.PipeWriter 和一个标记完成的函数
type PipeWriter struct {
	*io.PipeWriter
	done     func()
	buffered *bufio.Writer
}

// 重写读写方法
func (pw *PipeWriter) Write(p []byte) (n int, err error) {
	return pw.buffered.Write(p)
}

func (pr *PipeReader) Read(p []byte) (n int, err error) {
	return pr.buffered.Read(p)
}

func BufferedWaitPipe(bufSize int) (*PipeReader, *PipeWriter) {
	reader, writer := io.Pipe()
	var wg sync.WaitGroup
	wg.Add(1)

	//双端缓冲
	newWriter := bufio.NewWriterSize(writer, bufSize)
	newReader := bufio.NewReaderSize(reader, bufSize)

	return &PipeReader{
			PipeReader: reader,
			buffered:   newReader,
			wait:       wg.Wait,
		}, &PipeWriter{
			PipeWriter: writer,
			buffered:   newWriter,
			done: func() {
				newWriter.Flush()
				writer.Close()
				wg.Done()
			},
		}
}

func main() {
	// 定义常量：要传输的文件大小为1GB (1 * 1024 * 1024 * 1024字节)
	const fileSize = 1 * 1024 * 1024 * 1024 // 1GB

	// 定义管道缓冲区大小为64KB (64 * 1024字节)
	const bufferSize = 64 * 1024 // 64KB缓冲

	// 创建带缓冲的管道，返回读取端(pr)和写入端(pw)
	pr, pw := BufferedWaitPipe(bufferSize)

	// 启动一个goroutine来模拟写入1GB数据
	go func() {
		// 使用defer确保在函数退出时调用pw.done()
		// pw.done()会刷新缓冲区、关闭管道并通知等待组
		defer pw.done() // 会触发Flush和Close

		// 创建32KB的数据块用于写入
		data := make([]byte, 32*1024) // 32KB块

		// 用随机数据填充这个块(模拟真实文件内容)
		rand.Read(data) // 填充随机数据

		// 循环写入，直到累计写入量达到1GB
		for written := 0; written < fileSize; {
			// 将数据块写入管道
			n, err := pw.Write(data)
			if err != nil {
				log.Fatal("Write error:", err)
			}
			written += n // 累加已写入字节数
		}
	}()

	// 记录传输开始时间
	start := time.Now()

	// 创建SHA256哈希计算器
	hasher := sha256.New()

	// 从管道读取数据并计算哈希值
	// io.CopyBuffer使用64KB的缓冲区来提高读取效率
	if _, err := io.CopyBuffer(hasher, pr, make([]byte, 64*1024)); err != nil {
		log.Fatal("Read error:", err)
	}

	// 等待写入goroutine完成(通过pr.wait())
	pr.wait() // 等待写入完成

	// 计算传输耗时
	duration := time.Since(start)

	// 计算传输速度(MB/s)
	// fileSize转换为MB(除以1024*1024)，再除以耗时(秒)
	speed := float64(fileSize) / (1024 * 1024) / duration.Seconds()

	// 打印传输结果
	fmt.Printf("Transferred 1GB in %v, speed: %.2f MB/s\n", duration, speed)

	// 打印计算得到的SHA256校验值
	fmt.Printf("SHA256: %x\n", hasher.Sum(nil))
}
