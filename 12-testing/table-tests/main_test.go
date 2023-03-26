package main

import "testing"

func TestMySum(t *testing.T) { //pass

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3, 4}, 10},
		test{[]int{32, 10}, 42},
		test{[]int{40, 11, 10}, 61},
		test{[]int{3, 4, 5}, 12},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Excepted", v.answer, "Got", x)
		}
	}

}
