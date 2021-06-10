package middleware

import (
	"github.com/bcowtech/host"
	. "github.com/bcowtech/host-fasthttp/internal"
)

var _ host.Middleware = new(ErrorHandlerMiddleware)

type ErrorHandlerMiddleware struct {
	Handler ErrorHandler
}

func (m *ErrorHandlerMiddleware) Init(appCtx *host.AppContext) {
	var (
		fasthttphost = asFasthttpHost(appCtx.Host())
		preparer     = NewFasthttpHostPreparer(fasthttphost)
	)

	preparer.RegisterErrorHandler(m.Handler)
}
