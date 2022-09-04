package leetcode1935

import (
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	m := map[rune]uint8{}
	for _, v := range brokenLetters {
		m[v] = 0
	}
	count := 0
	words := strings.Split(text, " ")
	for _, val := range words {
		if contain(val, m) {
			count++
		}
	}
	return count
}

func contain(w string, m map[rune]uint8) bool {
	for _, v := range w {
		if _, ok := m[v]; ok {
			return false
		}
	}
	return true
}
