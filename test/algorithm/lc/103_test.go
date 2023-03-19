package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := make([]int, 0)
		q := queue
		queue = nil
		for i := 0; i < len(q); i++ {
			vals = append(vals, q[i].Val)
			if q[i].Left != nil {
				queue = append(queue, q[i].Left)
			}
			if q[i].Right != nil {
				queue = append(queue, q[i].Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			i, j := 0, len(vals)-1
			for i < j {
				vals[i], vals[j] = vals[j], vals[i]
				i++
				j--
			}
		}
		res = append(res, vals)
	}
	return res
}
