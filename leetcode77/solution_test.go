package leetcode77

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combine(t *testing.T) {
	expect := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	assert.ElementsMatch(t, expect, combine(4, 2))
}
