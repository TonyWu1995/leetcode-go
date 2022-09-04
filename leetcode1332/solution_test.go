package leetcode1332

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removePalindromeSub(t *testing.T) {
	assert.Equal(t, 1, removePalindromeSub("ababa"), "error")
	assert.Equal(t, 2, removePalindromeSub("abb"), "error")
}
