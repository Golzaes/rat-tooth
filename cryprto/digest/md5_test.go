package digest

import (
	"fmt"
	"testing"
)

func TestMD5HexLowercase(t *testing.T) {
	fmt.Println(Md5DigestHex(`asdsad`, false))
	fmt.Println(Md5DigestHex(`asdsad`, true))
	fmt.Println(Md5DigestBit32(`asdsad`, false))
	fmt.Println(Md5DigestBit32(`asdsad`, true))
}
