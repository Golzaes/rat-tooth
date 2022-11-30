package logs

import (
	"context"
	"fmt"
	"os"
	"sync"
)

// globalLogger is designed as a global logger in current process.
var global = &loggerAppliance{}

// loggerAppliance is the proxy of `Logger` to
// make logger change will affect all sub-logger.
type loggerAppliance struct {
	Logger
	mu sync.Mutex
}

func init() {
	global.SetLogger(DefaultLogger)
}

func (a *loggerAppliance) SetLogger(in Logger) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Logger = in
}

// SetLogger should be called before any other log call.
// And it is NOT THREAD SAFE.
func SetLogger(logger Logger) {
	global.SetLogger(logger)
}

// GetLogger returns global logger appliance as logger in current process.
func GetLogger() Logger {
	return global
}

// Log Print log by level and key values.
func Log(level Level, kvs ...interface{}) {
	_ = global.Log(level, kvs...)
}

// Context with context logger.
func Context(ctx context.Context) *Helper {
	return NewHelper(WithContext(ctx, global.Logger))
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	_ = global.Log(LevelDebug, DefaultMessageKey, fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	_ = global.Log(LevelDebug, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func Debugw(kvs ...interface{}) {
	_ = global.Log(LevelDebug, kvs...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	_ = global.Log(LevelInfo, DefaultMessageKey, fmt.Sprint(a...))
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	_ = global.Log(LevelInfo, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func Infow(kvs ...interface{}) {
	_ = global.Log(LevelInfo, kvs...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	_ = global.Log(LevelWarn, DefaultMessageKey, fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	_ = global.Log(LevelWarn, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func Warnw(kvs ...interface{}) {
	_ = global.Log(LevelWarn, kvs...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	_ = global.Log(LevelError, DefaultMessageKey, fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	_ = global.Log(LevelError, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func Errorw(kvs ...interface{}) {
	_ = global.Log(LevelError, kvs...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	_ = global.Log(LevelFatal, DefaultMessageKey, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	_ = global.Log(LevelFatal, DefaultMessageKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func Fatalw(kvs ...interface{}) {
	_ = global.Log(LevelFatal, kvs...)
	os.Exit(1)
}
