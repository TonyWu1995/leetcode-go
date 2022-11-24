package leetcode1828

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	assert.Equal(t, true, checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	assert.Equal(t, false, checkIfPangram("leetcode"))
}
