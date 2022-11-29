package fluent

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	log "github.com/golzaes/rat-tooth/observation/logs"
)

var _ log.Logger = (*Logger)(nil)

// Logger is fluent logger sdk.
type Logger struct {
	opts Fluent
	log  *fluent.Fluent
}

// FluentOption is fluentd logger option.
type FluentOption func(*Fluent)

type Fluent struct {
	timeout            time.Duration
	writeTimeout       time.Duration
	bufferLimit        int
	retryWait          int
	maxRetry           int
	maxRetryWait       int
	tagPrefix          string
	async              bool
	forceStopAsyncSend bool
}

// Log print the kv pairs log.
func (l *Logger) Log(level log.Level, kvs ...interface{}) error {
	if len(kvs) == 0 {
		return nil
	}
	if len(kvs)%2 != 0 {
		kvs = append(kvs, "KEYVALS UNPAIRED")
	}

	data := make(map[string]string, len(kvs)/2+1)

	for i := 0; i < len(kvs); i += 2 {
		data[fmt.Sprint(kvs[i])] = fmt.Sprint(kvs[i+1])
	}

	if err := l.log.Post(level.String(), data); err != nil {
		println(err)
	}
	return nil
}

func (l *Logger) Close() error {
	return l.log.Close()
}

// NewLogger new a std logger with options.
// target:
//
//	tcp://127.0.0.1:24224
//	unix://var/run/fluent/fluent.sock
func NewLogger(addr string, opts ...FluentOption) (*Logger, error) {
	fl := &Fluent{}
	fl.Init()
	for _, f := range opts {
		f(fl)
	}

	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	c := fluent.Config{
		Timeout:            fl.timeout,
		WriteTimeout:       fl.writeTimeout,
		BufferLimit:        fl.bufferLimit,
		RetryWait:          fl.retryWait,
		MaxRetry:           fl.maxRetry,
		MaxRetryWait:       fl.maxRetryWait,
		TagPrefix:          fl.tagPrefix,
		Async:              fl.async,
		ForceStopAsyncSend: fl.forceStopAsyncSend,
	}
	switch u.Scheme {
	case "tcp":
		host, port, err2 := net.SplitHostPort(u.Host)
		if err2 != nil {
			return nil, err2
		}
		if c.FluentPort, err = strconv.Atoi(port); err != nil {
			return nil, err
		}
		c.FluentNetwork = u.Scheme
		c.FluentHost = host
	case "unix":
		c.FluentNetwork = u.Scheme
		c.FluentSocketPath = u.Path
	default:
		return nil, fmt.Errorf("unknown network: %s", u.Scheme)
	}
	fle, err := fluent.New(c)
	if err != nil {
		return nil, err
	}
	return &Logger{
		opts: *fl,
		log:  fle,
	}, nil
}

func (f *Fluent) Init() {}

// WithTimeout allows you to set a custom timeout
func WithTimeout(timeout time.Duration) FluentOption {
	return func(f *Fluent) {
		f.timeout = timeout
	}
}

// WithWriteTimeout allows you to set a custom write timeout
func WithWriteTimeout(writeTimeout time.Duration) FluentOption {
	return func(f *Fluent) {
		f.writeTimeout = writeTimeout
	}
}

// WithBufferLimit allows you to set a custom buffer limit
func WithBufferLimit(bufferLimit int) FluentOption {
	return func(f *Fluent) {
		f.bufferLimit = bufferLimit
	}
}

// WithRetryWait allows you to set a custom retry wait
func WithRetryWait(retryWait int) FluentOption {
	return func(f *Fluent) {
		f.retryWait = retryWait
	}
}

// WithMaxRetry allows you to set a custom max retry count
func WithMaxRetry(maxRetry int) FluentOption {
	return func(f *Fluent) {
		f.maxRetry = maxRetry
	}
}

// WithMaxRetryWait allows you to set a custom max retry wait
func WithMaxRetryWait(maxRetryWait int) FluentOption {
	return func(f *Fluent) {
		f.maxRetryWait = maxRetryWait
	}
}

// WithTagPrefix allows you to set a custom tag prefix
func WithTagPrefix(tagPrefix string) FluentOption {
	return func(f *Fluent) {
		f.tagPrefix = tagPrefix
	}
}

// WithAsync allows you to set a custom async
func WithAsync(async bool) FluentOption {
	return func(f *Fluent) {
		f.async = async
	}
}

// WithForceStopAsyncSend allows you to set a custom forceStopAsyncSend
func WithForceStopAsyncSend(forceStopAsyncSend bool) FluentOption {
	return func(f *Fluent) {
		f.forceStopAsyncSend = forceStopAsyncSend
	}
}
