//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/nextmicro/logger"
	"next.bff.layout/internal/biz"
	"next.bff.layout/internal/conf"
	"next.bff.layout/internal/data"
	"next.bff.layout/internal/server"
	"next.bff.layout/internal/service"
	"next.bff.layout/internal/svc"

	"github.com/google/wire"
)

// wireApp init next application.
func wireApp(*conf.Data, *conf.Client, logger.Logger) (*Injector, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, svc.ProviderSet, InjectorSet, newApp))

	return new(Injector), nil, nil
}
