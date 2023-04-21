package leetcode290

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[string]string)
	splitString := strings.Split(s, " ")
	patternString := strings.Split(pattern, "")
	if len(splitString) != len(patternString) {
		return false
	}
	for i := 0; i < len(splitString); i++ {
		if _, ok := m[patternString[i]]; !ok {
			m[patternString[i]] = splitString[i]
			continue
		}
		_, ok := m[patternString[i]]
		if ok && m[patternString[i]] != splitString[i] {
			return false
		}
	}
	tmp := ""
	for _, v := range m {
		if strings.Contains(tmp, v) {
			return false
		}
		tmp += v
	}
	return true
}
