package server

import (
	"context"
	"net"
)

type Interceptor func(ctx context.Context, conn net.Conn, handler Handler) error

func chainInterceptors(interceptors []Interceptor) Interceptor {
	return func(ctx context.Context, conn net.Conn, handler Handler) error {
		return interceptors[0](ctx, conn, chainHandler(interceptors, 0, handler))
	}
}

func chainHandler(interceptors []Interceptor, curr int, finalHandler Handler) Handler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, conn net.Conn) error {
		return interceptors[curr+1](ctx, conn, chainHandler(interceptors, curr+1, finalHandler))
	}
}
