package leetcode217

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	assert.Equal(t, true, containsDuplicate([]int{1, 2, 3, 1}))
	assert.Equal(t, false, containsDuplicate([]int{1, 2, 3, 4}))
	assert.Equal(t, true, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
