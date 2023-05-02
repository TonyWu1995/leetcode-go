package leetcode784

import (
	"fmt"
	"unicode"
)

func letterCasePermutation(s string) []string {
	res := []string{s}
	for i := 0; i < len(s); i++ {
		length := len(res)
		if !unicode.IsDigit(rune(s[i])) {
			for j := 0; j < length; j++ {
				var newChar byte
				oldStr := res[j]
				if oldStr[i] >= 'a' && oldStr[i] <= 'z' {
					newChar = oldStr[i] - 32
				} else {
					newChar = oldStr[i] + 32
				}
				newStr := oldStr[:i] + fmt.Sprintf("%c", newChar) + oldStr[i+1:]
				res = append(res, newStr)
			}
		}
	}
	return res
}
