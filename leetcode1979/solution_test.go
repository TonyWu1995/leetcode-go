package leetcode1979

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindGCD(t *testing.T) {
	assert.Equal(t, 2, findGCD([]int{2, 5, 6, 9, 10}))
}
