package leetcode1603

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	big := 1
	medium := 1
	small := 1
	obj := Constructor(big, medium, small)

	if obj.slot[0] != 1 {
		t.Errorf("")
	}
	if obj.slot[1] != 1 {
		t.Errorf("")
	}
	if obj.slot[2] != 1 {
		t.Errorf("")
	}

}

func TestParkingSystem_AddCar(t *testing.T) {
	big := 1
	medium := 1
	small := 0
	
	obj := Constructor(big, medium, small)

	res := obj.AddCar(1)

	if res == false {
		t.Errorf("")
	}

	res = obj.AddCar(2)

	if res == false {
		t.Errorf("")
	}

	res = obj.AddCar(3)

	if res == true {
		t.Errorf("")
	}

}
