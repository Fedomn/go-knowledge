package lc

import (
	"math"
	"reflect"
	"testing"
)

func reverse(x int) int {
	val := x
	res := 0
	for val != 0 {
		// 注意这里的res每次可以先乘10，以达到位数的扩展，而不用事先计算好位数
		res = res*10 + val%10
		val = val / 10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	return res
}

func Test7(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{101, 101},
		{0, 0},
		{1534236469, 0},
	}

	for _, test := range tests {
		got := reverse(test.input)
		if !reflect.DeepEqual(test.expect, got) {
			t.Fatalf("expect: %v, got: %v", test.expect, got)
		}
	}
}
