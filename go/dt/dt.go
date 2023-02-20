package dt

import (
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MethodPair struct {
	Action     string
	Compensate string
	ProtoMsg   protoreflect.ProtoMessage
}

type TransactionActor interface {
	Execute(methodPairs ...MethodPair)
}

type transactionActor struct {
	dtmServerAddr string
	gid           string
}

func NewTransactionActor(addr string) TransactionActor {
	return &transactionActor{
		dtmServerAddr: addr,
		gid:           dtmgrpc.MustGenGid(addr),
	}
}

func (ta *transactionActor) withServerAction(action string) string {
	return ta.dtmServerAddr + action
}
func (ta *transactionActor) withServerCompensate(compensate string) string {
	return ta.dtmServerAddr + compensate
}

// Regist implements TransactionRegister
func (ta *transactionActor) Execute(methodPairs ...MethodPair) {
	saga := dtmgrpc.NewSagaGrpc(ta.dtmServerAddr, ta.gid)
	for _, methodPair := range methodPairs {
		saga = saga.Add(ta.withServerAction(methodPair.Action), ta.withServerCompensate(methodPair.Compensate), methodPair.ProtoMsg)
	}
	saga.WaitResult = true
	saga.Submit()
}

var _ TransactionActor = (*transactionActor)(nil)
