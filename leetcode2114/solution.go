package leetcode2114

import "strings"

func mostWordsFound(sentences []string) int {
	res := 0
	for _, sentence := range sentences {
		temp := len(strings.Split(sentence, " "))
		if temp > res {
			res = temp
		}
	}
	return res
}
