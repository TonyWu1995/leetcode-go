package leetcode2073

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_timeRequiredToBuy(t *testing.T) {

	assert.Equal(t, 6, timeRequiredToBuy([]int{2, 3, 2}, 2), "error")
	assert.Equal(t, 8, timeRequiredToBuy([]int{5, 1, 1, 1}, 0), "error")

}
