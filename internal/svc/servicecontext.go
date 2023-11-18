package svc

import (
	"context"

	v1 "github.com/nextmicro/next.data.layout/api/shorturl/v1"
	"github.com/nextmicro/next/registry"
	"github.com/nextmicro/next/transport/http"
	"next.bff.layout/internal/conf"
)

type ServiceContext struct {
	Config         *conf.Client
	ShortUrlClient v1.ShortUrlHTTPClient
}

//go:generate go run github.com/Bin-Huang/newc@v0.8.3
func NewServiceContext(c *conf.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

func (svc *ServiceContext) Run() error {
	shortUrlHTTPClient, err := NewGatewayClient(svc.Config)
	if err != nil {
		return err
	}

	svc.ShortUrlClient = shortUrlHTTPClient
	return err
}

func NewGatewayClient(c *conf.Client) (v1.ShortUrlHTTPClient, error) {
	client, err := http.NewClient(
		context.TODO(),
		http.WithConfig(c.GetGateway()),
		http.WithDiscovery(registry.DefaultRegistry))
	if err != nil {
		return nil, err
	}

	return v1.NewShortUrlHTTPClient(client), nil
}
