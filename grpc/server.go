package __

import (
	"context"
)

type MyServer struct{}

//	type AddServiceServer interface {
//		Add(context.Context, *AddRequest) (*AddReply, error)
//		mustEmbedUnimplementedAddServiceServer()
//	}
func (s *MyServer) Add(ctx context.Context, req *AddRequest) (*AddReply, error) {
	res := myAdd(req.A, req.B)
	return &AddReply{Res: res}, nil
}
func (s *MyServer) mustEmbedUnimplementedAddServiceServer() {

}

func myAdd(a int32, b int32) int32 {
	return a + b
}
