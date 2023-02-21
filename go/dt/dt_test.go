package dt

import (
	"testing"

	"github.com/diwangtseb/essayes/go/dt/pb"
)

const (
	dtm_server  = "localhost:36790"
	grpc_server = "localhost:1234"
)

func TestDtGrpc(t *testing.T) {
	actor := NewTransactionActor(dtm_server, grpc_server)

	actor.MsgExecute(MethodPair{
		Action:     "/dt.pb.TransService/TransIn",
		Compensate: "",
		ProtoMsg: &pb.TransInReq{
			Uid:    "1",
			Amount: "100",
		},
	})
}
