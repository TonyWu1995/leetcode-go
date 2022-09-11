package solution2391

func garbageCollection(garbage []string, travel []int) int {
	t := 0
	idx := len(travel) - 1
	for _, v := range travel {
		t += v
	}
	ch := make(chan int, 3)
	go calc(idx, t, 71, garbage, travel, ch)
	go calc(idx, t, 77, garbage, travel, ch)
	go calc(idx, t, 80, garbage, travel, ch)
	res := 0
	for i := 0; i < 3; i++ {
		res += <-ch
	}
	return res
}

func calc(idx int, t int, target int32, garbage []string, travel []int, ch chan int) {
	res := 0
	b := true
	for i := idx + 1; i >= 0; i-- {
		for _, v := range garbage[i] {
			if v == target {
				res++
			}
			if b && v == target {
				res += t
				b = false
			}
		}
		t -= travel[idx]
		if idx >= 1 {
			idx--
		}
	}
	ch <- res
}
