package leetcode1935

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canBeTypedWords(t *testing.T) {

	res := canBeTypedWords("hello world", "ad")
	assert.Equal(t, 1, res, "error")

	res = canBeTypedWords("leet code", "lt")
	assert.Equal(t, 1, res, "error")

	res = canBeTypedWords("leet code", "e")
	assert.Equal(t, 0, res, "error")
}
