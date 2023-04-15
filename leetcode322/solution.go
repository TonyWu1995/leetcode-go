package leetcode322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if coins == nil || len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = int(^uint(0) >> 1)
	}
	dp[0] = 0
	for _, c := range coins {
		for i := c; i < amount+1; i++ {
			if dp[i-c] != int(^uint(0)>>1) {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == int(^uint(0)>>1) {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
