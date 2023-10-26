package server

import (
	"github.com/nextmicro/logger"
	v1 "next.bff.layout/api/helloworld/v1"
	"next.bff.layout/internal/service"

	"github.com/nextmicro/next/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(greeter *service.GreeterService, logger logger.Logger) *grpc.Server {
	srv := grpc.NewServer()
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
