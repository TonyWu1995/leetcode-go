package leetcode228

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	assert.Equal(t, []string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	assert.Equal(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	assert.Equal(t, []string{"-1"}, summaryRanges([]int{-1}))
}
