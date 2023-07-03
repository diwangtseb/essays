package service

import (
	context "context"
	"fmt"

	"github.com/diwangtseb/essayes/go/dt/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.TransServiceServer = (*TransService)(nil)

type TransService struct {
	pb.UnimplementedTransServiceServer
}

// TransCancel implements pb.TransServiceServer.
func (*TransService) TransCancel(ctx context.Context, req *pb.TransCancelReq) (*pb.TransCancelReply, error) {
	fmt.Println(req.Uid + " [cancel] " + u32tostr(req.Amount))
	return &pb.TransCancelReply{
		Success: true,
	}, nil
}

// TransConfirm implements pb.TransServiceServer.
func (*TransService) TransConfirm(ctx context.Context, req *pb.TransConfirmReq) (*pb.TransConfirmReply, error) {
	fmt.Println(req.Uid + " [confirm] " + u32tostr(req.Amount))
	return &pb.TransConfirmReply{
		Success: true,
	}, nil
}

// TransTry implements pb.TransServiceServer.
func (*TransService) TransTry(ctx context.Context, req *pb.TransTryReq) (*pb.TransTryReply, error) {
	fmt.Println(req.Uid + " [try] " + u32tostr(req.Amount))
	return &pb.TransTryReply{
		Success: true,
	}, nil
}

// TransInRoll implements pb.TransServiceServer
func (ta *TransService) TransInRoll(ctx context.Context, req *pb.TransInReq) (*pb.TransInReply, error) {
	fmt.Println(req.Uid + " [decrease] " + u32tostr(req.Amount))
	return &pb.TransInReply{
		Success: true,
	}, nil
}

func u32tostr(u32 uint32) string {
	return fmt.Sprint(u32)
}

// TransIn implements pb.TransServiceServer
func (ts *TransService) TransIn(ctx context.Context, req *pb.TransInReq) (*pb.TransInReply, error) {
	fmt.Println(req.Uid + " [crease] " + u32tostr(req.Amount))
	if req.Uid == "3" {
		return &pb.TransInReply{
			Success: false,
		}, status.New(codes.Aborted, "test").Err()
	}
	return &pb.TransInReply{
		Success: true,
	}, nil
}

// TransOut implements pb.TransServiceServer
func (ts *TransService) TransOut(ctx context.Context, req *pb.TransOutReq) (*pb.TransOutReply, error) {
	fmt.Println(req.Uid + "decrease" + u32tostr(req.Amount))
	return &pb.TransOutReply{
		Success: true,
	}, nil
}
