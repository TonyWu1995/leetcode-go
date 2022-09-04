package leetcode2351

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repeatedCharacter(t *testing.T) {

	res := repeatedCharacter("abccbaacz")
	expect := byte('c')
	assert.Equal(t, expect, res, "error")
}
