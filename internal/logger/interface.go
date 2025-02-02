package logger

import "context"

type Logger interface {
	GetLevel() Level
	SetLevel(level Level)

	Debug(ctx context.Context, msg string)
	Debugf(ctx context.Context, msg string, args ...any)

	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, msg string, args ...any)

	Warn(ctx context.Context, msg string)
	Warnf(ctx context.Context, msg string, args ...any)

	Error(ctx context.Context, msg string)
	Errorf(ctx context.Context, msg string, args ...any)

	Panic(ctx context.Context, msg string)
	Panicf(ctx context.Context, msg string, args ...any)

	Fatal(ctx context.Context, msg string)               // Panic with os.Exit(1)
	Fatalf(ctx context.Context, msg string, args ...any) // Panic with os.Exit(1)
}
