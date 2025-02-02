package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/emikhalev/faraway_wow/internal/config"
	"github.com/emikhalev/faraway_wow/internal/logger"
	"github.com/emikhalev/faraway_wow/internal/tracer"
)

type Server struct {
	cfg config.Server

	listener net.Listener

	handler           Handler
	interceptorsChain Interceptor

	connectionsWg sync.WaitGroup

	stopErr chan error

	doneCh   chan struct{}
	doneOnce sync.Once
}

func New(cfg config.Server) *Server {
	return &Server{
		cfg:     cfg,
		stopErr: make(chan error),
		doneCh:  make(chan struct{}),
	}
}

func (srv *Server) Run(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", srv.cfg.Host, srv.cfg.Port)
	logger.Infof(ctx, "start listening on %s", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("cannot start listening on %s: %v", addr, err)
	}
	srv.listener = listener

	for {
		conn, err := listener.Accept()
		if err != nil {

			select {
			case <-srv.doneCh:
				logger.Infof(ctx, "TCP server shutting down...")
				return nil
			default:
				logger.Infof(ctx, "Error accepting connection: %v", err)
			}

		}

		srv.connectionsWg.Add(1)
		go func() {
			defer srv.connectionsWg.Done()
			srv.process(conn)
		}()

	}
}

func (srv *Server) Err() <-chan error {
	return srv.stopErr
}

func (srv *Server) Close() error {
	var err error
	srv.doneOnce.Do(func() {
		err = srv.listener.Close()
		close(srv.doneCh)
		srv.connectionsWg.Wait()
		srv.stopErr <- err
	})
	return err
}

func (srv *Server) process(conn net.Conn) {
	defer conn.Close()

	ctx := context.Background()
	ctx, span := tracer.Tracer().Start(ctx, "handleConnection")
	defer span.End()

	if srv.handler != nil {
		if srv.interceptorsChain != nil {
			if err := srv.interceptorsChain(ctx, conn, srv.handler); err != nil {
				logger.Errorf(ctx, "error processing TCP request: %v", err)
			}
			return
		}
		if err := srv.handler(ctx, conn); err != nil {
			logger.Errorf(ctx, "error processing TCP request: %v", err)
		}
		return
	}
}
