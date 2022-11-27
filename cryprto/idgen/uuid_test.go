package idgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var args = NewUuidArgs(2, `test`)
var GensUUid = []func() (string, error){
	args.GenerateUUID1,
	args.GenerateUUID3,
	args.GenerateUUID4,
	args.GenerateUUID5,
	args.GenerateUUID6,
	args.GenerateUUID7,
}

func TestUuiGen(t *testing.T) {
	for _, f := range GensUUid {
		s, err := f()
		require.Nil(t, err)
		require.True(t, IsValidUUID(s))
	}
}
