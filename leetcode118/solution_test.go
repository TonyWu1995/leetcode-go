package Leetcode120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generate(t *testing.T) {

	res := generate(3)
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}}, res, "error")
}
