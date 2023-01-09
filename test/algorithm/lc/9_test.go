package lc

import (
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res := 0
	val := x
	for val != 0 {
		res = res*10 + val%10
		val = val / 10
	}
	return res == x
}

func Test9(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
	}

	for _, test := range tests {
		got := isPalindrome(test.input)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.input)
		}
	}
}
