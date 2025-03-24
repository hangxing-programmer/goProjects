package __

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

}

func CreateServer() {
	s, _ := net.Listen("tcp", ":9999")
	service := MyServer{}
	grpcServer := grpc.NewServer()
	RegisterAddServiceServer(grpcServer, &service)
	grpcServer.Serve(s)
}

func CreateClient() {
	con, _ := grpc.Dial("localhost:9999", grpc.WithInsecure())
	defer con.Close()
	client := NewAddServiceClient(con)
	req := AddRequest{
		A: 8,
		B: 9,
	}
	res, err := client.Add(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Add: %v", err)
	}
	fmt.Println(res.Res)
}
