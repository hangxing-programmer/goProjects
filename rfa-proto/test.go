package rfa_proto

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

/*
*
proto编译命令:protoc --go_out=. --go-grpc_out=. *.proto
*/
func main() {
	conn, err := grpc.Dial("10.0.11.201:3501", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewFileServiceClient(conn)

	stream, err := c.ReadData(context.Background(), &DataReadRequest{
		VSID:       "CACHE_STORAGE1",
		LocationID: "test.txt",
		Offset:     0,
		Length:     -1, // -1表示读取整个文件
	})
	if err != nil {
		log.Fatalf("could not read data: %v", err)
	}

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive data: %v", err)
		}
		log.Printf("Received %d bytes of data.\n", chunk.SizeInBytes)

	}

	exists, err := c.Exists(context.Background(), &DataExistRequest{
		VSID:       "CACHE_STORAGE1",
		LocationID: "test.txt",
	})
	if err != nil {
		log.Fatalf("could not check: %v", err)
	}
	log.Printf("File exists: %v\n", exists.Exists)

	// 模拟写入数据块
	//data := []byte("Hello, world!")
	//if err := writeStream.Send(&DataChunk{
	//	Content:     data,
	//	SizeInBytes: int32(len(data)),
	//	IsLastChunk: true,
	//}); err != nil {
	//	log.Fatalf("could not send data: %v", err)
	//}
	//
	//handleResult, err = writeStream.CloseAndRecv()
	//if err != nil {
	//	log.Fatalf("could not finish writing: %v", err)
	//}
	//log.Printf("Write result: %d\n", handleResult.Handled)

	// 获取整个空间信息示例
	//spaceInfo, err := c.GetEntireSpaceInfo(context.Background(), &NullMsg{})
	//if err != nil {
	//	log.Fatalf("could not get space info: %v", err)
	//}
	//log.Printf("Total space: %d, Used space: %d, Free space: %d\n", spaceInfo.All, spaceInfo.Used, spaceInfo.Free)

}
