package leetcode1646

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaximumGenerated(t *testing.T) {
	assert.Equal(t, 3, getMaximumGenerated(7))
	assert.Equal(t, 1, getMaximumGenerated(2))
	assert.Equal(t, 2, getMaximumGenerated(3))
}
