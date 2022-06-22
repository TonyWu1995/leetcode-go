package leetcode1920

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	nums := []int{0, 2, 1, 5, 3, 4}

	res := buildArray(nums)
	expect := []int{0, 1, 2, 4, 5, 3}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("not match [0,1,2,4,5,3]")
	}

}
