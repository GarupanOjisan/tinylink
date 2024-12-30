package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EncodeBase62(t *testing.T) {
	assert.Equal(t, "Z", EncodeBase62(int64(61)))
}

func Test_DecodeBase62(t *testing.T) {
	assert.Equal(t, int64(61), DecodeBase62("Z"))
}
