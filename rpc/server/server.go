package server

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "ymir/rpc/proto"
)

type rpcServer struct {
	debug bool
}

func newRpcServer(debug bool) *grpc.Server {
	glog.Infof("new rpc server with debug %v", debug)
	s := grpc.NewServer()
	//TODO:
	//ymirpb.RegisterYmirServiceServer(s, &rpcServer{debug})
	return s
}