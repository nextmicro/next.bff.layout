package server

import (
	"github.com/nextmicro/logger"
	v1 "next.bff.layout/api/helloworld/v1"
	"next.bff.layout/internal/service"

	"github.com/nextmicro/next/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(greeter *service.GreeterService, logger logger.Logger) *http.Server {
	srv := http.NewServer()
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
