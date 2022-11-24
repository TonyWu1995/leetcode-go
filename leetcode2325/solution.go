package leetcode2265

import (
	"strings"
)

var alpha = [26]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z",
}

func decodeMessage(key string, message string) string {
	m := map[int32]string{}
	index := 0
	for _, v := range key {
		_, isExist := m[v]
		if v != 32 && !isExist {
			m[v] = alpha[index]
			index++
		}
	}
	res := strings.Builder{}
	for _, v := range message {
		if v == 32 {
			res.WriteString(" ")
		} else {
			res.WriteString(m[v])
		}
	}
	return res.String()
}
