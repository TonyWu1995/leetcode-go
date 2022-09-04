package leetcode1399

func countLargestGroup(n int) int {
	groups := [37]int{}
	max, countMax := 0, 0
	for i := 1; i <= n; i++ {
		s := mod10(i)
		groups[s]++
		if groups[s] == max {
			countMax++
		} else if groups[s] > max {
			max = groups[s]
			countMax = 1
		}
	}
	return countMax
}

func mod10(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return s
}
