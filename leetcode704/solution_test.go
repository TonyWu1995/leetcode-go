package leetcode704

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search_case1(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func Test_search_case2(t *testing.T) {
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
