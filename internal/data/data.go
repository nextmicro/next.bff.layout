package data

import (
	"github.com/nextmicro/logger"
	"next.bff.layout/internal/conf"
	"next.bff.layout/internal/svc"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewShortUrlRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	serviceContext *svc.ServiceContext
}

// NewData .
func NewData(c *conf.Data, logger logger.Logger, srv *svc.ServiceContext) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	return &Data{
		serviceContext: srv,
	}, cleanup, nil
}
