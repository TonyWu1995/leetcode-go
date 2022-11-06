package leetcode2221

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_triangularSum(t *testing.T) {
	assert.Equal(t, 8, triangularSum([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 5, triangularSum([]int{5}))
}
