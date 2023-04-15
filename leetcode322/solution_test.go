package leetcode322

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{
		1, 2, 5,
	}, 11))
}
