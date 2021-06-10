package fasthttp

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/host-fasthttp/internal"
)

func Startup(app interface{}) *host.Starter {
	var (
		starter = host.Startup(app)
	)

	host.RegisterHostService(starter, internal.FasthttpHostServiceInstance)

	return starter
}
