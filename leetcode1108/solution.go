package leetcode1108

import (
	"strings"
)

func defangIPaddr(address string) string {
	var sb strings.Builder
	for _, v := range address {
		if v == 46 {
			sb.WriteString("[.]")

		} else {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}
