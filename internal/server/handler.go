package server

import (
	"context"
	"net"
)

type Handler func(ctx context.Context, conn net.Conn) error
