package main

import (
	"calendar/config"
	"calendar/internal/handler"
	"calendar/internal/server"
	"calendar/internal/service"
	"calendar/internal/store"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("app error", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("config: %v", err)
	}

	memstore := store.NewMemStore()
	svc := service.NewEventService(memstore)

	mux := http.NewServeMux()
	handler.InitRoutes(mux, svc)

	srv, err := server.NewServer("localhost",
		server.WithPort(cfg.HTTP.Port),
		server.WithHandler(mux),
	)
	if err != nil {
		return err
	}

	var (
		stopCh = make(chan os.Signal)
		errCh  = make(chan error)
	)
	signal.Notify(stopCh, os.Interrupt)

	go func() {
		slog.Info("server is listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("http start: %v", err)
		}
	}()

	select {
	case <-stopCh:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("http shutdown: %v", err)
		}
	case err := <-errCh:
		return err
	}
	return nil
}
