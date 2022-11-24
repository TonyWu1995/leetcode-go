package leetcode1844

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_replaceDigits(t *testing.T) {
	assert.Equal(t, "abcdef", replaceDigits("a1c1e1"))
}
