package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "grpc/grpc"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (h *HelloServiceImpl) Hello(ctx context.Context, args *hello.StringReq) (*hello.StringRes, error) {
	reply := &hello.StringRes{
		Value: fmt.Sprintf("hello %s", args.GetValue()),
	}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", ":1314")
	if nil != err {
		log.Fatalln(err)
	}
	grpcServer.Serve(listen)
}
