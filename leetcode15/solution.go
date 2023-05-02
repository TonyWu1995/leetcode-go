package leetcode15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	m := make(map[[3]int][]int)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := [3]int{
					nums[i], nums[j], nums[k],
				}
				sort.Ints(triplet[:])
				if nums[i]+nums[j]+nums[k] == 0 {
					m[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	var result [][]int
	for _, val := range m {
		result = append(result, val)
	}
	return result
}
