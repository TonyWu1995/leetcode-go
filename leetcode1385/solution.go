package leetcode1385

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := 0; i < len(arr1); i++ {
		b := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				b = false
				break
			}
		}
		if b {
			res++
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
