package client

import (
	"context"
	"testing"

	"github.com/diwangtseb/essayes/go/dt/pb"
)

func TestClient(t *testing.T) {
	client := NewTransClient(":1234")
	client.TransIn(context.TODO(), &pb.TransInReq{
		Uid:    "1",
		Amount: 1,
	})
}
