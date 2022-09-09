package leetcode2073

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for {
		for i := range tickets {
			if tickets[k] == 0 {
				return res
			}
			if tickets[i] > 0 {
				tickets[i]--
				res++
			}
		}
	}
}
