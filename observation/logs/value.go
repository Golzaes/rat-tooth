package logs

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// DefaultCaller is a Valuer that returns the file and line.
var DefaultCaller = Caller(4)

// DefaultTimestamp is a Valuer that returns the current wallclock time.
var DefaultTimestamp = Timestamp(time.RFC3339)

// Valuer is returns a log value.
type Valuer func(ctx context.Context) any

// Value return the function value.
func Value(ctx context.Context, v any) any {
	if v, ok := v.(Valuer); ok {
		return v(ctx)
	}
	return v
}

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(context.Context) any {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return file[idx+1:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func(context.Context) any {
		return time.Now().Format(layout)
	}
}

func bindValues(ctx context.Context, kv []any) {
	for i := 1; i < len(kv); i += 2 {
		if v, ok := kv[i].(Valuer); ok {
			kv[i] = v(ctx)
		}
	}
}

func containsValuer(kv []any) bool {
	for i := 1; i < len(kv); i += 2 {
		if _, ok := kv[i].(Valuer); ok {
			return true
		}
	}
	return false
}
