package leetcode2433

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindArray(t *testing.T) {
	assert.Equal(t, []int{5, 7, 2, 3, 2}, findArray([]int{
		5, 2, 0, 3, 1,
	}))
}
