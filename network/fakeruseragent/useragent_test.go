package fakeruseragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomMobileUserAgent(t *testing.T) {
	assert.NotEmpty(t, RandomMobileUserAgent())
}

func TestRandomUserAgent(t *testing.T) {
	assert.NotEmpty(t, RandomUserAgent())
}
