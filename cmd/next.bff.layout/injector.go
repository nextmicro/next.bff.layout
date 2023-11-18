package main

import (
	"github.com/google/wire"
	"github.com/nextmicro/next"
	"next.bff.layout/internal/svc"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Next           *next.Next
	serviceContext *svc.ServiceContext
}
