package fakeruseragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakerUa = []func() string{
	RandomMobileUserAgent,
	RandomMobileUserAgent,
}

func TestRandomUserAgent(t *testing.T) {
	for _, f := range fakerUa {
		assert.NotEmpty(t, f())
	}
}
