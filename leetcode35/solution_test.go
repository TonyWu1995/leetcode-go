package leetcode35

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchInsert_case1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	assert.Equal(t, 2, searchInsert(nums, 5))
}

func Test_searchInsert_case2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	assert.Equal(t, 1, searchInsert(nums, 2))
}

func Test_searchInsert_case3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	assert.Equal(t, 4, searchInsert(nums, 7))
}
