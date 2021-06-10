package internal

import (
	"reflect"

	"github.com/bcowtech/host"
)

var _ host.HostService = new(FasthttpHostService)

type FasthttpHostService struct{}

func (p *FasthttpHostService) Init(h host.Host, app *host.AppContext) {
	if v, ok := h.(*FasthttpHost); ok {
		v.preInit()
	}
}

func (p *FasthttpHostService) InitComplete(h host.Host, app *host.AppContext) {
	if v, ok := h.(*FasthttpHost); ok {
		v.init()
	}
}

func (p *FasthttpHostService) GetHostType() reflect.Type {
	return typeOfHost
}
