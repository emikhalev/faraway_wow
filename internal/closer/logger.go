package closer

import (
	"context"
	"log"
)

type Logger interface {
	Error(ctx context.Context, msg string)
	Errorf(ctx context.Context, msg string, args ...any)

	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, msg string, args ...any)
}

type logger struct {
}

func newLogger() *logger {
	return new(logger)
}

func (l *logger) Error(_ context.Context, msg string) {
	log.Println(msg)
}

func (l *logger) Errorf(_ context.Context, msg string, args ...any) {
	log.Printf(msg, args...)
}

func (l *logger) Info(_ context.Context, msg string) {
	log.Println(msg)
}

func (l *logger) Infof(_ context.Context, msg string, args ...any) {
	log.Printf(msg, args...)
}
