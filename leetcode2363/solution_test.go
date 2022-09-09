package leetcode2363

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeSimilarItems(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 6},
		{3, 9},
		{4, 5},
	}, mergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}), "error")
}
