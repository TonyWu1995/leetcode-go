package leetcode1672

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sumCh := make(chan int)
		go sum(account, sumCh)
		temp := <-sumCh
		if temp > max {
			max = temp
		}
	}
	return max
}

func sum(account []int, channel chan int) {
	sum := 0
	for _, val := range account {
		sum += val
	}
	channel <- sum
}
