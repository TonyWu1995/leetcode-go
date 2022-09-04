package leetcode1396

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAverageTime(t *testing.T) {
	undergroundSystem := Constructor()

	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	res := undergroundSystem.GetAverageTime("Paradise", "Cambridge")
	assert.Equal(t, float64(14), res, "error")
}
