package hashing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// values given by https://md5calc.com/hash/md5/test

func TestHashMD4(t *testing.T) {
	h, err := NewHashingAlgorithm(HashMD4)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "db346d691d7acc4dc2625db19f9e3f52",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashMD5(t *testing.T) {
	h, err := NewHashingAlgorithm(HashMD5)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "098f6bcd4621d373cade4e832627b4f6",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha1(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA1)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha224(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA224)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha256(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA256)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha384(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA384)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha512(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA512)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha3_224(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA3_224)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "3797bf0afbbfca4a7bbba7602a2b552746876517a7f9b7ce2db0ae7b",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha3_256(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA3_256)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha3_384(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA3_384)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "e516dabb23b6e30026863543282780a3ae0dccf05551cf0295178d7ff0f1b41eecb9db3ff219007c4e097260d58621bd",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha3_512(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA3_512)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "9ece086e9bac491fac5c1d1046ca11d737b92a2b2ebd93f005d7b710110c0a678288166e7fbe796883a4f2e9b3ca9f484f521d0ce464345cc1aec96779149c14",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha512_224(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA512_224)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}

func TestHashSha512_256(t *testing.T) {
	h, err := NewHashingAlgorithm(HashSHA512_256)
	require.Nil(t, err)
	testCases := []struct {
		Input string
		Hash  string
	}{{
		Input: "test",
		Hash:  "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f",
	}}

	for _, testCase := range testCases {
		hash, err := h.Calculate(strings.NewReader(testCase.Input))
		require.Nil(t, err)
		assert.Equal(t, testCase.Hash, hash)
	}
}
