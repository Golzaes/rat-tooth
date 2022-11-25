package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"math"
	"strings"

	"github.com/OneOfOne/xxhash"
	"github.com/golzaes/rat-tooth/commonerrors"
	"github.com/golzaes/rat-tooth/reflection"
	str2 "github.com/golzaes/rat-tooth/strings"
	"github.com/spaolacci/murmur3"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// IHash defines a hashing algorithm.
type IHash interface {
	Calculate(reader io.Reader) (string, error)
	GetType() string
}

// referer src/crypto/crypto.go
const (
	HashMD4        = `MD4`        // import golang.org/x/crypto/md4
	HashMD5        = `MD5`        // import crypto/md5
	HashSHA1       = `SHA1`       // import crypto/sha1
	HashSHA224     = `SHA224`     // import crypto/sha256
	HashSHA256     = `SHA256`     // import crypto/sha256
	HashSHA384     = `SHA384`     // import crypto/sha512
	HashSHA512     = `SHA512`     // import crypto/sha512
	HashRIPEMD160  = `RIPEMD160`  // import golang.org/x/crypto/ripemd160
	HashSHA3_224   = `SHA3_224`   // import golang.org/x/crypto/sha3
	HashSHA3_256   = `SHA3_256`   // import golang.org/x/crypto/sha3
	HashSHA3_384   = `SHA3_384`   // import golang.org/x/crypto/sha3
	HashSHA3_512   = `SHA3_512`   // import golang.org/x/crypto/sha3
	HashSHA512_224 = `SHA512_224` // import crypto/sha512
	HashSHA512_256 = `SHA512_256` // import crypto/sha512
	HashMurmur     = `Murmur`     // import github.com/spaolacci/murmur3
	HashXXHash     = `xxhash`     // import github.com/OneOfOne/xxhash
)

type hashingAlgo struct {
	Hash hash.Hash
	Type string
}

func (h *hashingAlgo) Calculate(r io.Reader) (hashN string, err error) {
	if r == nil {
		return ``, commonerrors.ErrUndefined
	}

	if _, err := io.Copy(h.Hash, r); err != nil {
		return ``, err
	}

	hashN = hex.EncodeToString(h.Hash.Sum(nil))
	h.Hash.Reset()
	return
}

func (h *hashingAlgo) GetType() string {
	return h.Type
}

func NewHashingAlgorithm(hType string) (IHash, error) {
	var h hash.Hash
	switch hType {
	case HashMD4:
		h = md4.New()
	case HashMD5:
		h = md5.New()
	case HashSHA1:
		h = sha1.New()
	case HashSHA224:
		h = sha256.New224()
	case HashSHA256:
		h = sha256.New()
	case HashSHA384:
		h = sha512.New384()
	case HashSHA512:
		h = sha512.New()
	case HashRIPEMD160:
		h = ripemd160.New()
	case HashSHA3_224:
		h = sha3.New224()
	case HashSHA3_256:
		h = sha3.New256()
	case HashSHA3_384:
		h = sha3.New384()
	case HashSHA3_512:
		h = sha3.New512()
	case HashSHA512_224:
		h = sha512.New512_224()
	case HashSHA512_256:
		h = sha512.New512_256()
	case HashMurmur:
		h = murmur3.New64()
	case HashXXHash:
		h = xxhash.New64()
	}

	if h == nil {
		return nil, commonerrors.ErrNotFound
	}

	return &hashingAlgo{
		Hash: h,
		Type: hType,
	}, nil
}

func CalculateHash(text, htype string) string {
	hashing, err := NewHashingAlgorithm(htype)
	if err != nil {
		return ""
	}
	h, err := hashing.Calculate(strings.NewReader(text))
	if err != nil {
		return ""
	}
	return h
}

// HasLikelyHexHashStringEntropy states whether a string has an entropy
// which may entail it is a hexadecimal hash This is based on the work
// done by `detect-secrets` https://github.com/Yelp/detect-secrets/blob/2fc0e31f067af98d97ad0f507dac032c9506f667/detect_secrets/plugins/high_entropy_strings.py#L150
func HasLikelyHexHashStringEntropy(str string) bool {
	entropy := str2.CalculateStringShannonEntropy(str)
	entropy -= 1.2 / math.Log2(float64(len(str)))
	return entropy > 3.0
}

// IsLikelyHexHashString determines whether the string is likely to be a hexadecimal hash or not.
func IsLikelyHexHashString(str string) bool {
	if reflection.IsEmpty(str) {
		return false
	}
	if _, err := hex.DecodeString(str); err != nil {
		return false
	}
	return HasLikelyHexHashStringEntropy(str)
}
