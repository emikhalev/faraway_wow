package main

import (
	"context"
	"strings"
	"syscall"
	"time"

	"github.com/emikhalev/faraway_wow/internal/closer"
	"github.com/emikhalev/faraway_wow/internal/config"
	"github.com/emikhalev/faraway_wow/internal/interceptors/pow"
	"github.com/emikhalev/faraway_wow/internal/logger"
	"github.com/emikhalev/faraway_wow/internal/server"
	service "github.com/emikhalev/faraway_wow/internal/service"
)

const (
	closerDeadlineDefault = 5 // in seconds
)

func main() {
	ctx := context.Background()

	// config
	cfg := config.Get(ctx)

	// logger
	initLogger(cfg)

	// closer
	cl := closer.New(ctx,
		syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM).
		WithLogger(logger.DefaultLogger()).
		WithDeadline(closerDeadlineDefault * time.Second)

	// init service
	s := service.New()

	// run TCP server
	srv := server.New(cfg.Server).
		WithInterceptors(pow.Sha256).
		WithHandler(s.WoWHandler)

	cl.AddCloser(srv, "TCP Server")
	if err := srv.Run(ctx); err != nil {
		logger.Fatalf(ctx, "cannot start server: %v", err)
	}

	select {
	case err := <-srv.Err():
		logger.Errorf(ctx, "HTTP server stopped: %v", err)
	}
}

func initLogger(cfg config.Config) {
	switch strings.ToUpper(strings.TrimSpace(cfg.Logger.Level)) {
	case "DEBUG":
		logger.SetLevel(logger.DebugLevel)
	case "INFO":
		logger.SetLevel(logger.InfoLevel)
	case "WARN", "WARNING":
		logger.SetLevel(logger.WarnLevel)
	case "ERR", "ERROR":
		logger.SetLevel(logger.ErrorLevel)
	case "PANIC":
		logger.SetLevel(logger.PanicLevel)
	case "FATAL":
		logger.SetLevel(logger.FatalLevel)
	default:
		logger.SetLevel(logger.WarnLevel)
	}
}
