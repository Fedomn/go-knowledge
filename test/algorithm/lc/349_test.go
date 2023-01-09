package lc

func intersection(nums1 []int, nums2 []int) []int {
	vMap := make(map[int]bool)
	for _, n := range nums1 {
		vMap[n] = true
	}
	res := make([]int, 0)
	for _, n := range nums2 {
		isTrue := vMap[n]
		if isTrue {
			res = append(res, n)
			vMap[n] = false
		}
	}
	return res
}
