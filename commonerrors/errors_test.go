package commonerrors

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestAny(t *testing.T) {
	assert.True(t, Any(ErrNotImplemented, ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.False(t, Any(ErrNotImplemented, ErrInvalid, ErrUnknown))
	assert.True(t, Any(ErrNotImplemented, nil, ErrNotImplemented))
	assert.True(t, Any(nil, nil, ErrNotImplemented))
	assert.False(t, Any(ErrNotImplemented, nil, ErrInvalid, ErrUnknown))
	assert.False(t, Any(nil, ErrInvalid, ErrUnknown))
	assert.True(t, Any(fmt.Errorf("an error %w", ErrNotImplemented), ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.False(t, Any(fmt.Errorf("an error %w", ErrNotImplemented), ErrInvalid, ErrUnknown))
}

func TestNone(t *testing.T) {
	assert.False(t, None(ErrNotImplemented, ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.False(t, None(ErrNotImplemented, nil, ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.True(t, None(ErrNotImplemented, ErrInvalid, ErrUnknown))
	assert.True(t, None(ErrNotImplemented, nil, ErrInvalid, ErrUnknown))
	assert.True(t, None(nil, ErrInvalid, ErrUnknown))
	assert.False(t, None(nil, nil, ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.False(t, None(fmt.Errorf("an error %w", ErrNotImplemented), ErrInvalid, ErrNotImplemented, ErrUnknown))
	assert.True(t, None(fmt.Errorf("an error %w", ErrNotImplemented), ErrInvalid, ErrUnknown))
}

func TestContextErrorConversion(t *testing.T) {
	defer goleak.VerifyNone(t)
	task := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Asked to stop:", ctx.Err())
				return
			default:
				time.Sleep(time.Second * 1)
			}
		}
	}
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)
	err := ConvertContextError(cancelCtx.Err())
	require.NotNil(t, err)
	assert.True(t, Any(err, ErrTimeout, ErrCancelled))
}
