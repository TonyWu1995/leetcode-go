package Leetcode1207

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniqueOccurrences(t *testing.T) {

	assert.Equal(t, true, uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}), "error")

}
