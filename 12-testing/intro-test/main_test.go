package main

import "testing"

func TestMySum(t *testing.T) { //pass
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMySumm(t *testing.T) { //fail
	x := mySumm(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
