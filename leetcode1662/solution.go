package leetcode1828

import (
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.EqualFold(strings.Join(word1, ""), strings.Join(word2, ""))
}
