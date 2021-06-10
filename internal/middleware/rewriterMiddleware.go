package middleware

import (
	"github.com/bcowtech/host"
	. "github.com/bcowtech/host-fasthttp/internal"
)

type RewriterMiddleware struct {
	Handler RewriteHandler
}

func (m *RewriterMiddleware) Init(appCtx *host.AppContext) {
	var (
		fasthttphost = asFasthttpHost(appCtx.Host())
		preparer     = NewFasthttpHostPreparer(fasthttphost)
	)

	preparer.RegisterRewriteHandler(m.Handler)
}
