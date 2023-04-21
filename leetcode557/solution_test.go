package leetcode557

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
}
