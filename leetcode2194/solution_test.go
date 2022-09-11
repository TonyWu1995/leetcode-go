package leetcode2194

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	assert.Equal(t, []string{"K1", "K2", "L1", "L2"},
		cellsInRange("K1:L2"))
	assert.Equal(t, []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		cellsInRange("A1:F1"))
}
