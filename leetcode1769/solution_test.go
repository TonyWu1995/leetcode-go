package leetcode1769

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minOperations(t *testing.T) {
	assert.Equal(t, []int{1, 1, 3}, minOperations("110"))
	assert.Equal(t, []int{11, 8, 5, 4, 3, 4}, minOperations("001011"))
}
