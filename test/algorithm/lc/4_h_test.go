package lc

import (
	"fmt"
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
	"reflect"
	"testing"
)

// log -> 二分搜索
// 根据中位数的定义，
// 当 m+n 是奇数时，中位数是两个有序数组中的第 (m+n)/2 个元素 (注意计算时 idx = (m+n)/2 + 1)
// 当 m+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2 个元素和第 (m+n)/2+1 个元素的平均值
// 因此，这道题可以转化成寻找两个有序数组中的第 k 小的数，其中 k 为 (m+n)/2 或 (m+n)/2+1
//
// refer 最后一个视频: https://leetcode.cn/problems/median-of-two-sorted-arrays/solution/di-k-xiao-shu-jie-fa-ni-zhen-de-dong-ma-by-geek-8m/
// 通过二分的方法，每次比较两个数组的 k/2 元素大小(因为两个数组的 k/2 元素个数之和就是 k，因此比较后，剔除的数肯定小于第 k 个)；在剔除的小边的数后，进行下一轮，同时 k = k - 剔除的个数
//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		// 比如：totalLen=3时，这里的k=2，代表取第2个数。但在 getKthElem 方法里需要将k-1(1)才能取到真正的第2个数
		k := totalLen/2 + 1
		return float64(getKthElem(nums1, nums2, k))
	} else {
		// 比如：totalLen=4时，这里k1=2,k2=3，代表去第2和第3个数
		k1 := totalLen / 2
		k2 := totalLen/2 + 1
		return float64(getKthElem(nums1, nums2, k1)+getKthElem(nums1, nums2, k2)) / 2
	}
}

// 这里的 k 代表第 k 个数，而不是 idx，因此需要将 k-1 才是 idx，因此下面在取数组元素时，都是 k-1
func getKthElem(nums1, nums2 []int, k int) int {
	n1StartIdx, n2StartIdx := 0, 0
	for {
		if n1StartIdx == len(nums1) {
			fmt.Println("return, n1StartIdx,", n1StartIdx, k, nums2[n2StartIdx+k-1])
			// nums1 已经扫描完，则返回 nums2 的第 k 个元素
			return nums2[n2StartIdx+k-1]
		}
		if n2StartIdx == len(nums2) {
			fmt.Println("return, n2StartIdx,", n2StartIdx, k, nums1[n1StartIdx+k-1])
			return nums1[n1StartIdx+k-1]
		}
		if k == 1 {
			fmt.Println("return, k=1,", n1StartIdx, n2StartIdx)
			// 找到第 1 小的元素
			return Min(nums1[n1StartIdx], nums2[n2StartIdx])
		}

		half := k / 2
		// 准备下一轮的起始 idx
		n1StartIdxNew := Min(n1StartIdx+half, len(nums1)) - 1
		n2StartIdxNew := Min(n2StartIdx+half, len(nums2)) - 1
		fmt.Println("start", n1StartIdxNew, n2StartIdxNew, ",", nums1[n1StartIdxNew], nums2[n2StartIdxNew])
		if nums1[n1StartIdxNew] <= nums2[n2StartIdxNew] {
			k -= n1StartIdxNew - n1StartIdx + 1
			n1StartIdx = n1StartIdxNew + 1
		} else {
			k -= n2StartIdxNew - n2StartIdx + 1
			n2StartIdx = n2StartIdxNew + 1
		}
	}
}

func Test4(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.expect, findMedianSortedArrays(test.nums1, test.nums2)) {
			t.Fatalf("expect: %v, got: %v", test.expect, findMedianSortedArrays(test.nums1, test.nums2))
		}
	}
}
