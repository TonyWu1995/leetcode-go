package leetcode2079

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wateringPlants(t *testing.T) {
	assert.Equal(t, 14, wateringPlants([]int{2, 2, 3, 3}, 5))
	assert.Equal(t, 30, wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4))
	assert.Equal(t, 49, wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8))
	assert.Equal(t, 17, wateringPlants([]int{3, 2, 4, 2, 1}, 6))
}
