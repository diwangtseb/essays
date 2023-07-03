package dt

import (
	"context"
	"log"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MethodPair struct {
	Action     string
	Compensate string
	Try        string
	Confirm    string
	Cancel     string
	ProtoMsg   protoreflect.ProtoMessage
}

type TransactionActor interface {
	ExecuteSaga(ctx context.Context, methodPairs ...MethodPair)
	ExecuteMsg(ctx context.Context, methodPairs ...MethodPair)
	ExecuteXa(ctx context.Context, methodPairs ...MethodPair)
	ExecuteTcc(ctx context.Context, methodPairs ...MethodPair)
}

type transactionActor struct {
	dtmServerAddr  string
	grpcServerAddr string
	gid            string
	logger         log.Logger
}

func NewTransactionActor(dtmAddr, grpcAddr string) TransactionActor {
	return &transactionActor{
		dtmServerAddr:  dtmAddr,
		grpcServerAddr: grpcAddr,
		gid:            dtmgrpc.MustGenGid(dtmAddr),
		logger:         *log.Default(),
	}
}

func (ta *transactionActor) withServerAction(action string) string {
	return ta.grpcServerAddr + action
}
func (ta *transactionActor) withServerCompensate(compensate string) string {
	return ta.grpcServerAddr + compensate
}

// MsgExecute implements TransactionActor
func (ta *transactionActor) ExecuteMsg(ctx context.Context, methodPairs ...MethodPair) {
	msg := dtmgrpc.NewMsgGrpc(ta.dtmServerAddr, ta.gid)
	for _, v := range methodPairs {
		msg = msg.Add(ta.withServerAction(v.Action), v.ProtoMsg)
	}
	msg.WaitResult = true
	if err := msg.Submit(); err != nil {
		ta.logger.Fatal(err)
	}
}

// Regist implements TransactionRegister
func (ta *transactionActor) ExecuteSaga(ctx context.Context, methodPairs ...MethodPair) {
	saga := dtmgrpc.NewSagaGrpc(ta.dtmServerAddr, ta.gid)

	for _, methodPair := range methodPairs {
		saga = saga.Add(ta.withServerAction(methodPair.Action), ta.withServerCompensate(methodPair.Compensate), methodPair.ProtoMsg)
	}
	saga.WaitResult = true
	if err := saga.Submit(); err != nil {
		ta.logger.Fatal(err)
	}
}

// ExecuteXa implements TransactionActor.
func (ta *transactionActor) ExecuteXa(ctx context.Context, methodPairs ...MethodPair) {
	empty := &emptypb.Empty{}
	err := dtmgrpc.XaGlobalTransaction(ta.dtmServerAddr, ta.gid, func(xa *dtmgrpc.XaGrpc) error {
		for _, methodPair := range methodPairs {
			err := xa.CallBranch(methodPair.ProtoMsg, ta.withServerAction(methodPair.Action), empty)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		ta.logger.Fatal(err)
	}
}

// ExecuteTcc implements TransactionActor.
func (ta *transactionActor) ExecuteTcc(ctx context.Context, methodPairs ...MethodPair) {
	empty := &emptypb.Empty{}
	err := dtmgrpc.TccGlobalTransaction(ta.dtmServerAddr, ta.gid, func(tcc *dtmgrpc.TccGrpc) error {
		for _, methodPair := range methodPairs {
			err := tcc.CallBranch(methodPair.ProtoMsg, ta.withServerAction(methodPair.Try), ta.withServerAction(methodPair.Confirm), ta.withServerAction(methodPair.Cancel), empty)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		ta.logger.Fatal(err)
	}
}

var _ TransactionActor = (*transactionActor)(nil)
