package leetcode15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	assert.Contains(t, [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	},
		threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//assert.Equal(t, [][]int{},
	//	threeSum([]int{0,1,1}))
}
