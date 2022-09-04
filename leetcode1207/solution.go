package Leetcode1207

func uniqueOccurrences(arr []int) bool {
	countArr := map[int]int{}
	for _, val := range arr {
		countArr[val]++
	}
	countArr2 := map[int]int{}
	for _, val := range countArr {
		countArr2[val]++
		if countArr2[val] != 1 {
			return false
		}

	}
	return true
}
