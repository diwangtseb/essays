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

func TestDtGrpc(t *testing.T) {
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
