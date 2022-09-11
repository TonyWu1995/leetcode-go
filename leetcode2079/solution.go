package leetcode2079

func wateringPlants(plants []int, capacity int) int {
	step := 0
	currentWater := capacity
	for i := 0; i < len(plants); i++ {
		if currentWater < plants[i] {
			step += i * 2
			currentWater = capacity
		}
		step++
		currentWater = currentWater - plants[i]
	}
	return step
}
