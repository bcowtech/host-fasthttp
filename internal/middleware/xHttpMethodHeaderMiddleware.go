package middleware

import (
	"github.com/bcowtech/host"
	. "github.com/bcowtech/host-fasthttp/internal"
)

var _ host.Middleware = new(XHttpMethodHeaderMiddleware)

type XHttpMethodHeaderMiddleware struct {
	Headers []string
}

func (m *XHttpMethodHeaderMiddleware) Init(appCtx *host.AppContext) {
	var (
		fasthttphost = asFasthttpHost(appCtx.Host())
		preparer     = NewFasthttpHostPreparer(fasthttphost)
	)

	routeResolveModule := &XHttpMethodHeaderRouteResolveModule{
		headers: m.Headers,
	}
	preparer.RegisterRouteResolveModule(routeResolveModule)
}
