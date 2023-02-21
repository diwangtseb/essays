package service

import (
	context "context"
	"fmt"

	"github.com/diwangtseb/essayes/go/dt/pb"
)

var _ pb.TransServiceServer = (*TransService)(nil)

type TransService struct {
	pb.UnimplementedTransServiceServer
}

// TransIn implements pb.TransServiceServer
func (ts *TransService) TransIn(ctx context.Context, req *pb.TransInReq) (*pb.TransInReply, error) {
	fmt.Println(req.Uid + "增加" + req.Amount)
	return &pb.TransInReply{
		Success: true,
	}, nil
}

// TransOut implements pb.TransServiceServer
func (ts *TransService) TransOut(ctx context.Context, req *pb.TransOutReq) (*pb.TransOutReply, error) {
	fmt.Println(req.Uid + "减少" + req.Amount)
	return &pb.TransOutReply{
		Success: true,
	}, nil
}
