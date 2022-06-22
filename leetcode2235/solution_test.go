package main

import "testing"

func Test_sum(t *testing.T) {
	num1 := 12
	num2 := 5
	res := sum(num1, num2)

	if res != 17 {
		t.Errorf("not match 17")
	}
}
