package leetcode1528

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_restoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	assert.Equal(t, "leetcode", restoreString(s, indices))
}
