package leetcode1769

func minOperations(boxes string) []int {
	size := len(boxes)
	res := make([]int, size)
	for i, ops, cnt := 0, 0, 0; i < size; i++ {
		res[i] += ops
		if boxes[i] == 49 {
			cnt += 1
		} else {
			cnt += 0
		}
		ops += cnt
	}
	for i, ops, cnt := size-1, 0, 0; i >= 0; i-- {
		res[i] += ops
		if boxes[i] == 49 {
			cnt += 1
		} else {
			cnt += 0
		}
		ops += cnt
	}
	return res
}
func minOperations_2(boxes string) []int {
	size := len(boxes)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		ans := 0
		for j := 0; j < size; j++ {
			if boxes[j] == 49 {
				if i < j {
					ans += j - i
				} else {
					ans += i - j
				}
			}
		}
		res[i] = ans
	}
	return res
}
