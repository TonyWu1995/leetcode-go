package leetcode228

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var res []string
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 { //!continue number
			res = append(res, buildResString(nums[start], nums[i-1]))
			start = i
		}
		if len(nums)-1 == i {
			res = append(res, buildResString(nums[start], nums[i]))
		}
	}
	return res
}

func buildResString(i1, i2 int) string {
	if i1 == i2 {
		return strconv.Itoa(i1)
	}
	return strconv.Itoa(i1) + "->" + strconv.Itoa(i2)
}
