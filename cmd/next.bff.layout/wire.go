//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/nextmicro/logger"
	"github.com/nextmicro/next"
	"next.bff.layout/internal/biz"
	"next.bff.layout/internal/conf"
	"next.bff.layout/internal/data"
	"next.bff.layout/internal/server"
	"next.bff.layout/internal/service"

	"github.com/google/wire"
)

// wireApp init next application.
func wireApp(*conf.Data, *conf.Client, logger.Logger) (*next.Next, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
