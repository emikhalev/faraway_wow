//nolint:gochecknoglobals
package logger

import "context"

var (
	defaultLogger Logger
)

//nolint:gochecknoinits
func init() {
	defaultLogger = NewBasic(WarnLevel)
}

func DefaultLogger() Logger {
	return defaultLogger
}

func GetLevel() Level {
	return defaultLogger.GetLevel()
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

func Debug(ctx context.Context, msg string) {
	defaultLogger.Debug(ctx, msg)
}

func Debugf(ctx context.Context, msg string, args ...any) {
	defaultLogger.Debugf(ctx, msg, args...)
}

func Info(ctx context.Context, msg string) {
	defaultLogger.Info(ctx, msg)
}

func Infof(ctx context.Context, msg string, args ...any) {
	defaultLogger.Infof(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string) {
	defaultLogger.Warn(ctx, msg)
}

func Warnf(ctx context.Context, msg string, args ...any) {
	defaultLogger.Warnf(ctx, msg, args...)
}

func Error(ctx context.Context, msg string) {
	defaultLogger.Error(ctx, msg)
}

func Errorf(ctx context.Context, msg string, args ...any) {
	defaultLogger.Errorf(ctx, msg, args...)
}

func Panic(ctx context.Context, msg string) {
	defaultLogger.Panic(ctx, msg)
}

func Panicf(ctx context.Context, msg string, args ...any) {
	defaultLogger.Panicf(ctx, msg, args...)
}

func Fatal(ctx context.Context, msg string) {
	defaultLogger.Fatal(ctx, msg)
}

func Fatalf(ctx context.Context, msg string, args ...any) {
	defaultLogger.Fatalf(ctx, msg, args...)
}
