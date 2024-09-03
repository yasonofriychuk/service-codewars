package main

import (
	"net"

	"google.golang.org/grpc"

	codewars "github.com/yasonofriychuk/service-codewars/internal/generated/proto/service"
	"github.com/yasonofriychuk/service-codewars/internal/grpc/service"
)

func main() {
	gRPCServer := grpc.NewServer()
	codewars.RegisterCodewarsServer(gRPCServer, service.New())

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// Запускаем обработчик gRPC-сообщений
	if err := gRPCServer.Serve(l); err != nil {
		panic(err)
	}
}
