package digest

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type md5Gen struct {
}

func md5Digest(str string) string {
	m5 := md5.New()
	m5.Write([]byte(str))
	return hex.EncodeToString(m5.Sum(nil))
}

func Md5DigestBit32(str string, isUpper bool) string {
	m := md5Digest(str)
	if !isUpper {
		return m
	}
	return strings.ToUpper(m)
}

func Md5DigestHex(str string, isUpper bool) string {
	m := md5Digest(str)[8:24]
	if !isUpper {
		return m
	}
	return strings.ToUpper(m)
}

func md5DigestSalt(str, salt string, isUpper bool) string {
	m5 := md5.New()
	m5.Write([]byte(str))

	return hex.EncodeToString(m5.Sum(nil))
}
