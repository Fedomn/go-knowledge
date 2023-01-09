package lc

import (
	"reflect"
	"testing"
)

var vMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := vMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+letters[i])
		}
	}
}

func Test17(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.expect, letterCombinations(test.input)) {
			t.Fatal("expected:", test.expect, "actual:", letterCombinations(test.input))
		}
	}
}
