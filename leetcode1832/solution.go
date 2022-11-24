package leetcode1828

import (
	"strings"
)

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	for i := 'a'; i <= 'z'; i++ {
		if strings.Index(sentence, string(i)) == -1 {
			return false
		}
	}
	return true
}
