package response

import http "github.com/valyala/fasthttp"

const (
	// response name in RequestCtx user store
	RESPONSE_INVARIANT_NAME string = "github.com/bcowtech/host-fasthttp/response::Response"
)

func Success(ctx *http.RequestCtx, contentType string, body []byte) {
	ctx.Success(contentType, body)

	storeResponse(
		ctx,
		&responseImpl{
			flag:        SUCCESS,
			statusCode:  ctx.Response.StatusCode(),
			contentType: contentType,
			body:        body,
		},
	)
}

func Failure(ctx *http.RequestCtx, contentType string, message []byte, statusCode int) {
	ctx.SetStatusCode(statusCode)
	ctx.Success(contentType, message)

	storeResponse(
		ctx,
		&responseImpl{
			flag:        FAILURE,
			statusCode:  statusCode,
			contentType: contentType,
			body:        message,
		},
	)
}

func storeResponse(ctx *http.RequestCtx, resp Response) {
	ctx.SetUserValue(RESPONSE_INVARIANT_NAME, resp)
}