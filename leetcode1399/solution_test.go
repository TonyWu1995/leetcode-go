package leetcode1399

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countLargestGroup(t *testing.T) {
	input := 13
	res := countLargestGroup(input)
	assert.Equal(t, 4, res, "error")
}

func Test_countLargestGroup_case2(t *testing.T) {
	input := 2
	res := countLargestGroup(input)
	assert.Equal(t, 2, res, "error")
}
