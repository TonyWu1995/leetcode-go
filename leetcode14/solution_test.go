package leetcode14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
