package Leetcode120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getRow(t *testing.T) {

	res := getRow(3)
	assert.Equal(t, []int{1, 3, 3, 1}, res, "error")
}
