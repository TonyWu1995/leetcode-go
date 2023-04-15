package leetcode1979

func findGCD(nums []int) int {
	maxVal := 0
	minVal := 99999999
	for _, v := range nums {
		maxVal = max(maxVal, v)
		minVal = min(minVal, v)
	}
	return gcd(maxVal, minVal)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
