package leetcode557

import (
	"strings"
)

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	res := ""
	for _, v := range split {
		res = res + reverse(v) + " "
	}
	return strings.Trim(res, " ")
}

func reverse(s string) string {
	var str string
	for _, v := range s {
		str = string(v) + str
	}
	return str
}
