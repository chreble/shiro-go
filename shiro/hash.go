package shiro

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

var (
	algoList = []string{"md5", "sha1", "sha256", "sha512"}
	algoName = "md5"
	itr      = 1
	enc      = base64.StdEncoding
)

type SimpleHash struct {
	algo    string
	encType string // Encoding type: hex , base64
	iter    uint8  // no of hash iterations
	Data    []byte
	salt    []byte
}

type HashedValue []byte

func SetAlgoName(name string) error {
	for _, val := range algoList {
		if val == name {
			algoName = name
			return nil
		}
	}
	return ErrUnknownAlgorithm
}

func SetIteration(i int) {
	itr = i
}

// This function will generate the Hash of input data
func GenerateHash(data string) HashedValue {
	var hashVal HashedValue
	switch algoName {
	case "md5":
		hashVal = md5Val(data)
	case "sha1":
		hashVal = sha1Val(data)
	case "sha256":
		hashVal = sha256Val(data)
	case "sha512":
		hashVal = sha512Val(data)
	}
	return hashVal
}

func md5Val(data string) HashedValue {
	h := md5.New()
	io.WriteString(h, data)
	out := fmt.Sprintf("%s", h.Sum(nil))
	return []byte(out)
}

func sha1Val(data string) HashedValue {
	h := sha1.New()
	io.WriteString(h, data)
	out := fmt.Sprintf("%s", h.Sum(nil))
	return []byte(out)
}

func sha256Val(data string) HashedValue {
	h := sha256.New()
	io.WriteString(h, data)
	out := fmt.Sprintf("%s", h.Sum(nil))
	return []byte(out)
}

func sha512Val(data string) HashedValue {
	h := sha512.New()
	io.WriteString(h, data)
	out := fmt.Sprintf("%s", h.Sum(nil))
	return []byte(out)
}

func (h HashedValue) Hex() HashedValue {
	str := fmt.Sprintf("%x", h)
	return []byte(str)
}

func (h HashedValue) Base64() HashedValue {
	dst := make([]byte, enc.EncodedLen(len(h)))
	enc.Encode(dst, h)
	return dst
}
