package zap

import (
	"fmt"

	log "github.com/golzaes/rat-tooth/observation/logs"
	"go.uber.org/zap"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *zap.Logger
}

func NewLogger(z *zap.Logger) *Logger {
	return &Logger{z}
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}

func (l *Logger) Close() error {
	return l.Sync()
}

func (l *Logger) Log(level log.Level, kvs ...interface{}) error {
	if len(kvs) == 0 || len(kvs)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", kvs))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(kvs); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(kvs[i]), kvs[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}
