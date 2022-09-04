package Leetcode120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal(t *testing.T) {

	//[[2],[3,4],[6,5,7],[4,1,8,3]]
	res := minimumTotal([][]int{
		{
			2,
		},
		{
			3, 4,
		},
		{
			6, 5, 7,
		}, {
			4, 1, 8, 3,
		},
	})
	assert.Equal(t, 11, res, "error")

}
