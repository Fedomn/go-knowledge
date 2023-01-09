package lc

import (
	"reflect"
	"testing"
)

// 滑动窗口
// 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
// 一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
// 以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求。
func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	//resStr := ""
	indexes := make(map[byte]int)
	for right < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		indexes[s[right]] = right
		right++
		if res < right-left {
			res = right - left
			//resStr = s[left:right]
		}
	}
	//fmt.Println(resStr, right, left)
	return res
}

func Test3(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"a", 1},
		{"aab", 2},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.expected, lengthOfLongestSubstring(test.s)) {
			t.Fatalf("expected: %v, but got: %v", test.expected, lengthOfLongestSubstring(test.s))
		}
	}
}
