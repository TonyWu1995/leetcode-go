package leetcode1561

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxCoins(t *testing.T) {
	assert.Equal(t, 18, maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))

	assert.Equal(t, 4, maxCoins([]int{2, 4, 5}))
	assert.Equal(t, 9, maxCoins([]int{2, 4, 1, 2, 7, 8}))
}
