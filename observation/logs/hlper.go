package logs

import (
	"context"
	"fmt"
	"os"
)

// DefaultMessageKey default message key.
var DefaultMessageKey = "msg"

type HelperOption func(*Helper)

// Helper is a logger helper.
type Helper struct {
	logger Logger
	msgKey string
}

// NewHelper new a logger helper.
func NewHelper(logger Logger, opts ...HelperOption) *Helper {
	h := &Helper{}
	h.Init()
	h.logger = logger
	for _, f := range opts {
		f(h)
	}
	return h
}

func (h *Helper) Init() {
	h.msgKey = DefaultMessageKey
}

// WithContext returns a shallow copy of h with its context changed
// to ctx. The provided ctx must be non-nil.
func (h *Helper) WithContext(ctx context.Context) *Helper {
	return &Helper{
		msgKey: h.msgKey,
		logger: WithContext(ctx, h.logger),
	}
}

// Log Print log by level and keyvals.
func (h *Helper) Log(level Level, kvs ...interface{}) {
	_ = h.logger.Log(level, kvs...)
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	_ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	_ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(kvs ...interface{}) {
	_ = h.logger.Log(LevelDebug, kvs...)
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	_ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	_ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func (h *Helper) Infow(kvs ...interface{}) {
	_ = h.logger.Log(LevelInfo, kvs...)
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	_ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	_ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(kvs ...interface{}) {
	_ = h.logger.Log(LevelWarn, kvs...)
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	_ = h.logger.Log(LevelError, h.msgKey, fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	_ = h.logger.Log(LevelError, h.msgKey, fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(kvs ...interface{}) {
	_ = h.logger.Log(LevelError, kvs...)
}

// Fatal logs a message at fatal level.
func (h *Helper) Fatal(a ...interface{}) {
	_ = h.logger.Log(LevelFatal, h.msgKey, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func (h *Helper) Fatalf(format string, a ...interface{}) {
	_ = h.logger.Log(LevelFatal, h.msgKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func (h *Helper) Fatalw(kvs ...interface{}) {
	_ = h.logger.Log(LevelFatal, kvs...)
	os.Exit(1)
}

// WithMessageKey with message key.
func WithMessageKey(k string) HelperOption {
	return func(h *Helper) {
		h.msgKey = k
	}
}
