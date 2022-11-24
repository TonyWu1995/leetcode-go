package leetcode2265

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decodeMessage(t *testing.T) {
	res := decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")
	fmt.Println(res)
	assert.Equal(t, "this is a secret", res)
}
