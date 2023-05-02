package leetcode46

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	intput := []int{1, 2, 3}
	expect := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	assert.ElementsMatch(t, expect, permute(intput))
}
