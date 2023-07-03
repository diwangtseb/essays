package dt

import (
	"context"
	"testing"

	"github.com/diwangtseb/essayes/go/dt/pb"
)

const (
	dtm_server  = "localhost:36790"
	grpc_server = "localhost:1234"
)

func TestDtGrpcSaga(t *testing.T) {
	actor := NewTransactionActor(dtm_server, grpc_server)

	actor.ExecuteSaga(context.TODO(), MethodPair{
		Action:     "/dt.pb.TransService/TransIn",
		Compensate: "/dt.pb.TransService/TransInRoll",
		ProtoMsg: &pb.TransInReq{
			Uid:    "1",
			Amount: 100,
		},
	})
}

func TestDtGrpcMsg(t *testing.T) {
	actor := NewTransactionActor(dtm_server, grpc_server)

	actor.ExecuteMsg(context.TODO(), MethodPair{
		Action:     "/dt.pb.TransService/TransIn",
		Compensate: "/dt.pb.TransService/TransInRoll",
		ProtoMsg: &pb.TransInReq{
			Uid:    "1",
			Amount: 100,
		},
	})
}

func TestDtGrpcXa(t *testing.T) {
	actor := NewTransactionActor(dtm_server, grpc_server)

	actor.ExecuteXa(context.TODO(), MethodPair{
		Action:     "/dt.pb.TransService/TransIn",
		Compensate: "/dt.pb.TransService/TransInRoll",
		ProtoMsg: &pb.TransInReq{
			Uid:    "1",
			Amount: 100,
		},
	})
}

func TestDtGrpcTcc(t *testing.T) {
	actor := NewTransactionActor(dtm_server, grpc_server)

	actor.ExecuteTcc(context.TODO(), MethodPair{
		Try:      "/dt.pb.TransService/TransTry",
		Confirm:  "/dt.pb.TransService/TransConfirm",
		Cancel:   "/dt.pb.TransService/TransCancel",
		ProtoMsg: &pb.TransInReq{Uid: "1", Amount: 100},
	})
}
