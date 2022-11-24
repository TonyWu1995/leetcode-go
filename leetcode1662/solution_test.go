package leetcode1828

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayStringsAreEqual(t *testing.T) {
	assert.Equal(t, true, arrayStringsAreEqual([]string{
		"ab", "c",
	}, []string{
		"a", "bc",
	}))
}
