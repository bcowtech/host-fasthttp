package middleware

import (
	"runtime/debug"

	. "github.com/bcowtech/host-fasthttp/internal"
	"github.com/bcowtech/host-fasthttp/response"
)

var _ RequestHandleModule = new(LoggingHandleModule)

type LoggingHandleModule struct {
	successor      RequestHandleModule
	loggingService LoggingService
}

func (h *LoggingHandleModule) CanSetSuccessor() bool {
	return true
}

func (h *LoggingHandleModule) SetSuccessor(successor RequestHandleModule) {
	h.successor = successor
}

func (h *LoggingHandleModule) ProcessRequest(ctx *RequestCtx) {
	if h.successor != nil {
		eventLog := h.loggingService.CreateEventLog()

		defer func() {
			err := recover()
			if err != nil {
				resp := h.getResponse(ctx)

				defer func() {
					if resp != nil {
						eventLog.WriteResponse(ctx)
					} else {
						eventLog.WriteError(ctx, err, debug.Stack())
					}
				}()

				// NOTE: we should not handle error here, due to the underlying RequestHandler
				// will handle it.
			} else {
				eventLog.WriteResponse(ctx)
			}
			eventLog.Flush()
		}()
		h.successor.ProcessRequest(ctx)
	}
}

func (h *LoggingHandleModule) getResponse(ctx *RequestCtx) response.Response {
	obj := ctx.UserValue(response.RESPONSE_INVARIANT_NAME)
	v, ok := obj.(response.Response)
	if ok {
		return v
	}
	return nil
}
