package leetcode784

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	expect := []string{
		"a1b2", "A1b2", "A1B2", "a1B2",
	}
	res := letterCasePermutation("a1b2")
	assert.ElementsMatch(t, expect, res)
}
