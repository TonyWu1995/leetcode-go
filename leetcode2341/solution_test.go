package leetcode2341

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	assert.Equal(t, []int{3, 1}, numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}), "error")
}
