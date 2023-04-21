package leetcode290

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordPattern(t *testing.T) {
	assert.Equal(t, true, wordPattern("abba", "dog cat cat dog"))
	assert.Equal(t, false, wordPattern("aaaa", "dog cat cat dog"))
	assert.Equal(t, false, wordPattern("abba", "dog cat cat fish"))
	assert.Equal(t, true, wordPattern("syys", "a abc abc a"))
	assert.Equal(t, false, wordPattern("abba", "dog dog dog dog"))
}
