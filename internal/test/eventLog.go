package test

import (
	"fmt"

	fasthttp "github.com/bcowtech/host-fasthttp"
)

var _ fasthttp.EventLog = new(EventLog)

type EventLog struct{}

func (l *EventLog) WriteError(ctx *fasthttp.RequestCtx, err interface{}, stackTrace []byte) {
	fmt.Println("EventLog.WriteError()")
}
func (l *EventLog) WriteResponse(ctx *fasthttp.RequestCtx) {
	fmt.Println("EventLog.WriteResponse()")
}
func (l *EventLog) Flush() {
	fmt.Println("EventLog.Flush()")
}
