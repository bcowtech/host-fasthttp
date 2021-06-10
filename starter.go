package fasthttp

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/host-fasthttp/internal"
)

func Startup(app interface{}, middlewares ...host.Middleware) *host.Starter {
	var (
		starter = host.Startup(app, middlewares...)
	)

	host.RegisterHostService(starter, internal.FasthttpHostServiceInstance)

	return starter
}
