package server

import (
	"net/http"
	"strconv"
)

const (
	defaultHTTPPort = 8080
)

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == 0 {
		port = defaultHTTPPort
	} else {
		port = options.port
	}

	var handler http.Handler
	if options.handler == nil {
		handler = http.NewServeMux()
	} else {
		handler = options.handler
	}

	return &http.Server{
		Addr:    addr + ":" + strconv.Itoa(port),
		Handler: handler,
	}, nil
}
