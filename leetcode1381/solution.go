package leetcode1381

type CustomStack struct {
	position int
	data     []int
	size     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		position: 0,
		data:     make([]int, maxSize),
		size:     maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.position == this.size {
		return
	}
	this.data[this.position] = x
	this.position++
}

func (this *CustomStack) Pop() int {
	if this.position == 0 {
		return -1
	}
	this.position--
	return this.data[this.position]
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < this.position; i++ {
		this.data[i] += val
	}
}
