package server

import (
	"errors"
	"net/http"
)

type options struct {
	port    int
	handler http.Handler
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = port
		return nil
	}
}

func WithHandler(handler http.Handler) Option {
	return func(options *options) error {
		options.handler = handler
		return nil
	}
}
