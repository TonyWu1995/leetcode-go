package leetcode1561

import (
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	round := len(piles)
	sum := 0
	for i := round / 3; i < round-1; i += 2 {
		sum += piles[i]
	}
	return sum
}
