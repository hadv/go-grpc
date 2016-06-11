package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	echo "github.com/hadv/grpc/echo"
	"google.golang.org/grpc"
)

func Run() error {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	echo.RegisterEchoServiceServer(server, newEchoServer())
	server.Serve(listen)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
