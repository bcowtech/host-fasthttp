package internal

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/bcowtech/host"
	"github.com/valyala/fasthttp"
)

var _ host.Host = new(FasthttpHost)

type FasthttpHost struct {
	Server         *Server
	ListenAddress  string
	EnableCompress bool
	Version        string

	requestWorker *FasthttpRequestWorker

	initialized bool
	running     bool
	wg          sync.WaitGroup
}

func (h *FasthttpHost) Start(ctx context.Context) {
	if !h.initialized {
		panic("the FasthttpHost havn't be initialized yet")
	}
	if h.running {
		return
	}
	h.running = true

	defer func() {
		err := recover()
		if err != nil {
			h.running = false
		}
	}()

	s := h.Server

	log.Printf("[bcowtech/host-fasthttp] %s listening on address %s\n", h.Server.Name, h.ListenAddress)
	if err := s.ListenAndServe(h.ListenAddress); err != nil {
		log.Fatalf("[bcowtech/host-fasthttp] %% error in ListenAndServe: %v\n", err)
	}
}

func (h *FasthttpHost) Stop(ctx context.Context) error {
	h.wg.Wait()

	var (
		server = h.Server
	)

	defer func() {
		h.running = false
	}()

	return server.Shutdown()
}

func (h *FasthttpHost) preInit() {
	h.requestWorker = NewFasthttpRequestWorker()
}

func (h *FasthttpHost) init() {
	if h.initialized {
		return
	}

	if h.Server == nil {
		h.Server = &Server{}
	}

	h.requestWorker.init()
	h.configRequestHandler()
	h.configCompress()
	h.configListenAddress()

	h.initialized = true
}

func (h *FasthttpHost) configRequestHandler() {
	s := h.Server
	var requestHandler RequestHandler

	if s.Handler != nil {
		requestHandler = s.Handler
	} else if h.requestWorker != nil {
		requestHandler = h.requestWorker.ProcessRequest
	}

	s.Handler = func(ctx *RequestCtx) {
		h.wg.Add(1)
		defer func() {
			h.wg.Done()
		}()
		requestHandler(ctx)
	}
}

func (h *FasthttpHost) configCompress() {
	s := h.Server
	if h.EnableCompress {
		s.Handler = fasthttp.CompressHandler(s.Handler)
	}
}

func (h *FasthttpHost) configListenAddress() {
	host, port, err := splitHostPort(h.ListenAddress)
	if err != nil {
		panic(err)
	}

	if len(port) == 0 {
		port = DEFAULT_HTTP_PORT
	}
	h.ListenAddress = net.JoinHostPort(host, port)
}
