package main

import (
	"github.com/golang/glog"
	echo "github.com/hadv/grpc/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() echo.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *echo.Message) (*echo.Message, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *echo.Message) (*echo.Message, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}
