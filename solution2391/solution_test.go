package solution2391

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_garbageCollection(t *testing.T) {
	assert.Equal(t, 21, garbageCollection([]string{"G", "P", "GP", "GG"}, []int{
		2, 4, 3,
	}))

	//assert.Equal(t, 37, garbageCollection([]string{"MMM", "PGM", "GP"}, []int{
	//	3, 10,
	//}))
}
