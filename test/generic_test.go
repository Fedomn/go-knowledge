// refer to https://go.dev/doc/tutorial/generics

package main

import (
	"fmt"
	"testing"
)

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64 as types for map values.
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestGeneric(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[int]float64{
		1: 35.98,
		2: 26.99,
	}
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
