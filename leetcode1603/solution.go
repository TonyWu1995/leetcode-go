package leetcode1603

type ParkingSystem struct {
	slot [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	index := carType - 1
	if this.slot[index] > 0 {
		this.slot[index]--
		return true
	}

	return false
}
