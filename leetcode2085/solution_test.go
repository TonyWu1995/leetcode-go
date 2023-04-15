package leetcode2085

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountWords(t *testing.T) {
	res := countWords([]string{
		"leetcode", "is", "amazing", "as", "is",
	}, []string{
		"amazing", "leetcode", "is",
	})
	assert.Equal(t, 2, res)
	////["a","ab"]
	////["a","a","a","ab"]
	res = countWords([]string{
		"a", "ab",
	}, []string{
		"a", "a", "a", "ab",
	})
	assert.Equal(t, 1, res)
}
