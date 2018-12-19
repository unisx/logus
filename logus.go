package logus

import "go.uber.org/zap"

var logger *zap.Logger

type (
	Config = zap.Config
	Option = zap.Option
	Field  = zap.Field
)

func init() {
	InDevelopment(zap.Development(), zap.AddCaller(), zap.AddCallerSkip(1))
}

func newDevelopmentConfig() Config {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	return config
}

func newProductionConfig() Config {
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	return config
}

// InDevelopment is a reasonable development logging configuration.
// Logging is enabled at DebugLevel and above.
//
// It enables development mode (which makes DPanicLevel logs panic), uses a
// console encoder, writes to standard error, and disables sampling.
// StackTraces are automatically included on logs of WarnLevel and above.
func InDevelopment(opts ...Option) {
	var err error
	logger, err = newDevelopmentConfig().Build(opts...)
	if err != nil {
		panic(err)
	}
}

// InProduction is a reasonable production logging configuration.
// Logging is enabled at InfoLevel and above.
//
// It uses a JSON encoder, writes to standard error, and enables sampling.
// StackTraces are automatically included on logs of ErrorLevel and above.
func InProduction(opts ...Option) {
	var err error
	logger, err = newProductionConfig().Build(opts...)
	if err != nil {
		panic(err)
	}
}

// WithOptions applies the supplied Options.
func WithOptions(opts ...Option) {
	logger = logger.WithOptions(opts...)
}

// WithCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option).
func WithCallerSkip(skip int) {
	logger = logger.WithOptions(zap.AddCallerSkip(skip))
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func With(fields ...Field) {
	logger = logger.With(fields...)
}

// Named adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func Named(s string) {
	logger = logger.Named(s)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...Field) {
	logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...Field) {
	logger.DPanic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() error {
	return logger.Sync()
}
