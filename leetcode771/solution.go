package leetcode771

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	result := 0
	for _, char := range jewels {
		result += strings.Count(stones, string(char))
	}
	return result
}
