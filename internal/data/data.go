package data

import (
	"context"

	"github.com/nextmicro/logger"
	v1 "github.com/nextmicro/next.data.layout/api/shorturl/v1"
	"github.com/nextmicro/next/registry"
	"github.com/nextmicro/next/transport/http"
	"next.bff.layout/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGatewayClient, NewGreeterRepo, NewShortUrlRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	shortUrl v1.ShortUrlHTTPClient
}

// NewData .
func NewData(c *conf.Data, logger logger.Logger, shortUrl v1.ShortUrlHTTPClient) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	return &Data{
		shortUrl: shortUrl,
	}, cleanup, nil
}

func NewGatewayClient(c *conf.Client) (v1.ShortUrlHTTPClient, error) {
	client, err := http.NewClient(
		context.TODO(),
		c.GetGateway(),
		http.WithDiscovery(registry.DefaultRegistry))
	if err != nil {
		return nil, err
	}

	return v1.NewShortUrlHTTPClient(client), nil
}
