package leetcode1282

import (
	"fmt"
	"testing"
)

func Test_groupThePeople_case1(t *testing.T) {
	groupSize := []int{3, 3, 3, 3, 3, 1, 3}

	res := groupThePeople(groupSize)

	expect := [][]int{
		{5},
		{0, 1, 2},
		{3, 4, 6},
	}

	if len(res) == 0 {
		t.Errorf("not match")
	}

	fmt.Println(expect)

}

func Test_groupThePeople_case2(t *testing.T) {
	groupSize := []int{2, 1, 3, 3, 3, 2}

	res := groupThePeople(groupSize)

	expect := [][]int{
		{1},
		{0, 5},
		{2, 3, 4},
	}

	if len(res) == 0 {
		t.Errorf("not match")
	}
	fmt.Println(expect)
}
