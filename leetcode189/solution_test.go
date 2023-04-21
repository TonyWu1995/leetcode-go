package leetcode189

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, input)
}
