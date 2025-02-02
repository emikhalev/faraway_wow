package logger

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.opentelemetry.io/otel/trace"

	"github.com/rs/zerolog"
)

type basicLogger struct {
	level Level

	muLogger sync.Mutex
	logger   zerolog.Logger
}

func NewBasic(level Level) Logger {
	logger := &basicLogger{
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
	logger.SetLevel(level)

	return logger
}

func (l *basicLogger) GetLevel() Level {
	return l.level
}

func (l *basicLogger) SetLevel(level Level) {
	l.muLogger.Lock()
	defer l.muLogger.Unlock()

	l.logger = l.logger.Level(l.mapLevelToZeroLogLevel(level))
	l.level = level
}

func (l *basicLogger) Debug(ctx context.Context, msg string) {
	if l.level > DebugLevel {
		return
	}

	l.logger.
		Debug().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Debugf(ctx context.Context, msg string, args ...any) {
	if l.level > DebugLevel {
		return
	}

	l.logger.
		Debug().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) Info(ctx context.Context, msg string) {
	if l.level > InfoLevel {
		return
	}

	l.logger.
		Info().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Infof(ctx context.Context, msg string, args ...any) {
	if l.level > InfoLevel {
		return
	}

	l.logger.
		Info().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) Warn(ctx context.Context, msg string) {
	if l.level > WarnLevel {
		return
	}

	l.logger.
		Warn().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Warnf(ctx context.Context, msg string, args ...any) {
	if l.level > WarnLevel {
		return
	}

	l.logger.
		Warn().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) Error(ctx context.Context, msg string) {
	if l.level > ErrorLevel {
		return
	}

	l.logger.
		Error().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Errorf(ctx context.Context, msg string, args ...any) {
	if l.level > ErrorLevel {
		return
	}

	l.logger.
		Error().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) Panic(ctx context.Context, msg string) {
	if l.level > PanicLevel {
		return
	}

	l.logger.
		Panic().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Panicf(ctx context.Context, msg string, args ...any) {
	if l.level > PanicLevel {
		return
	}

	l.logger.
		Panic().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) Fatal(ctx context.Context, msg string) {
	if l.level > FatalLevel {
		return
	}

	l.logger.
		Fatal().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(msg)
}

func (l *basicLogger) Fatalf(ctx context.Context, msg string, args ...any) {
	if l.level > FatalLevel {
		return
	}

	l.logger.
		Fatal().
		Fields(l.getSpanFieldsContext(ctx)).
		Msg(fmt.Sprintf(msg, args...))
}

func (l *basicLogger) getSpanFieldsContext(ctx context.Context) map[string]any {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()
	spanID := span.SpanContext().SpanID().String()
	return map[string]any{
		"trace-id": traceID,
		"span-id":  spanID,
	}
}

func (l *basicLogger) mapLevelToZeroLogLevel(level Level) zerolog.Level {
	switch level {
	case DebugLevel:
		return zerolog.DebugLevel
	case InfoLevel:
		return zerolog.InfoLevel
	case WarnLevel:
		return zerolog.WarnLevel
	case ErrorLevel:
		return zerolog.ErrorLevel
	case PanicLevel:
		return zerolog.PanicLevel
	case FatalLevel:
		return zerolog.FatalLevel
	}
	return zerolog.Disabled
}
