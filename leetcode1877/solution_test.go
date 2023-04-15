package leetcode1877

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPairSum(t *testing.T) {
	assert.Equal(t, 7, minPairSum([]int{3, 5, 2, 3}))
}
