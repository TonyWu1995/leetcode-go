package leetcode1108

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_defangIPaddr(t *testing.T) {
	address := "1.1.1.1"
	assert.Equal(t, "1[.]1[.]1[.]1", defangIPaddr(address))
}
