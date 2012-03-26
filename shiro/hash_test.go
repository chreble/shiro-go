package shiro

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

var (
	in  = "This is test string"
	out = ""
)

func init() {
	algoName = "sha1"
	h := sha1.New()
	io.WriteString(h, in)
	out = fmt.Sprintf("%s", h.Sum(nil))
}

func TestHashedValue(t *testing.T) {
	val := GenerateHash(in)
	if string(val) == out {
		t.Log("%s Equals %s", out, string(val))
	} else {
		t.Fatal("Expected %s Got %s", out, val)
	}
}
