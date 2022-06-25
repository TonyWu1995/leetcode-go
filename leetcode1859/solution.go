package leetcode1859

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	res := make([]string, len(words))
	for _, word := range words {
		val, _ := strconv.Atoi(word[len(word)-1:])
		res[val-1] = word[:len(word)-1]

	}
	return strings.Join(res, " ")
}
