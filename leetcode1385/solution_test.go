package leetcode1385

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheDistanceValue(t *testing.T) {
	assert.Equal(t, 2, findTheDistanceValue(
		[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2,
	))
}
