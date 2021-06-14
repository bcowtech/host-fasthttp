package internal

import (
	"reflect"

	"github.com/valyala/fasthttp"
)

const (
	DEFAULT_HTTP_PORT   = "80"
	HEADER_XHTTP_METHOD = "X-Http-Method"

	StatusNotFound = 404
)

var (
	FasthttpHostServiceInstance = new(FasthttpHostService)

	typeOfHost = reflect.TypeOf(FasthttpHost{})
)

// import
type (
	Server         = fasthttp.Server
	RequestHandler = fasthttp.RequestHandler
	RequestCtx     = fasthttp.RequestCtx
)

// interface
type (
	RouteResolver interface {
		ResolveHttpMethod(ctx *RequestCtx) string
		ResolveHttpPath(ctx *RequestCtx) string
	}

	RouteResolveModule interface {
		RouteResolver

		CanSetSuccessor() bool
		SetSuccessor(successor RouteResolver)
	}

	RequestHandleModule interface {
		CanSetSuccessor() bool
		SetSuccessor(successor RequestHandleModule)
		ProcessRequest(ctx *RequestCtx, recover *RecoverService)
	}
)

// function
type (
	ErrorHandler   func(ctx *RequestCtx, err interface{})
	RewriteHandler func(ctx *RequestCtx, path *RoutePath) *RoutePath
)
