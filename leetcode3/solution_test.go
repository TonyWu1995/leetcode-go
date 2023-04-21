package leetcode3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	//assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
}
