package leetcode2011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_finalValueAfterOperations(t *testing.T) {

	res := finalValueAfterOperations([]string{"--X", "X++", "X++"})
	assert.Equal(t, 1, res)
}
