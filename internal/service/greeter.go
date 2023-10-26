package service

import (
	"context"

	v1 "next.bff.layout/api/helloworld/v1"
	"next.bff.layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc         *biz.GreeterUsecase
	shortUrlUc *biz.ShortUrlUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, shortUrlUc *biz.ShortUrlUsecase) *GreeterService {
	return &GreeterService{uc: uc, shortUrlUc: shortUrlUc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.shortUrlUc.Expand(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g}, nil
}
