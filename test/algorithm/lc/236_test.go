package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
// 有3种情况：
// 1. p 和 q 位于 root 的左右两侧
// 2. p == root，并且 q 在 root 子树中
// 3. q == root，并且 p 在 root 子树中
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
