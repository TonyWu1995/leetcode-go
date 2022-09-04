package leetcode2352

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_equalPairs(t *testing.T) {

	res := equalPairs([][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	})

	assert.Equal(t, 1, res, "error")
}
