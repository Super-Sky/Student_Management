package util

import (
	"testing"
)

func TestTest(t *testing.T) {
	i := int64(0)
	SetInt64(&i)
	println(i)
}
