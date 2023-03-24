package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

// 二叉树的右视图: 记录每一层的最后一个元素
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, level int) {
		if n != nil {
			if len(res) == level {
				res = append(res, n.Val)
			}
			// 优先找右边
			dfs(n.Right, level+1)
			// 如果右边没有，则选择左边
			dfs(n.Left, level+1)
		}
		return
	}
	dfs(root, 0)
	return res
}
