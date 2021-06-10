package middleware

import (
	"reflect"

	"github.com/bcowtech/host-fasthttp/internal"
)

var (
	typeOfHost           = reflect.TypeOf(internal.FasthttpHost{})
	typeOfRequestHandler = reflect.TypeOf(internal.RequestHandler(nil))

	TAG_URL = "url"
)

type (
	LoggingService interface {
		CreateEventLog() EventLog
	}

	EventLog interface {
		WriteError(ctx *internal.RequestCtx, err interface{}, stackTrace []byte)
		WriteResponse(ctx *internal.RequestCtx)
		Flush()
	}
)
