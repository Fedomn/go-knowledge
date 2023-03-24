package util

import "fmt"

// https://www.wikiwand.com/zh/%E6%A0%91%E7%9A%84%E9%81%8D%E5%8E%86
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 可以将遍历节点的顺序想象成：fmt为root，left!=nil为left，right!=nil为right
// 前序遍历: root->left->right
func PreOrder(root *TreeNode) {
	fmt.Println(root.Val)
	if root.Left != nil {
		PreOrder(root.Left)
	}
	if root.Right != nil {
		PreOrder(root.Right)
	}
}

// 中序遍历：left->root->right
// 注意：这个顺序也针对于子树，也就是说：访问左子树或者右子树的时候我们按照同样的方式遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
func InOrder(root *TreeNode) {
	if root.Left != nil {
		InOrder(root.Left)
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		InOrder(root.Right)
	}
}

// 中序遍历：栈方法
func InOrder2(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		// 一直取left节点，直到最后一个节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈的root代表最后一个left节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 打印left值
		fmt.Println(root.Val)
		// 取右边的节点递归下去
		root = root.Right
	}
}

// 后序遍历：left->right->root
func PostOrder(root *TreeNode) {
	if root.Left != nil {
		PostOrder(root.Left)
	}
	if root.Right != nil {
		PostOrder(root.Right)
	}
	fmt.Println(root.Val)
}

// 二叉树的最大深度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(height(root.Left), height(root.Right)) + 1
}

// 利用二叉树的level遍历
// 由于每层只有一个数，因此，len(res) == level 时，代表遇到当前层
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

// 二叉树最大宽度
// 类似堆存数组时的index一样，即：root编号为idx时候，左孩子编号: 2*idx，右孩子变化: 2*idx+1
func widthOfBinaryTree(root *TreeNode) int {
	levelMin := map[int]int{}
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, depth, index int) int {
		if node == nil {
			return 0
		}
		if _, ok := levelMin[depth]; !ok {
			levelMin[depth] = index // 每一层最先访问到的节点会是最左边的节点，即每一层编号的最小值
		}
		leftChildWidth := dfs(node.Left, depth+1, index*2)
		rightChildWidth := dfs(node.Right, depth+1, index*2+1)
		currentLevelWidth := index - levelMin[depth]
		return Max(currentLevelWidth+1, Max(leftChildWidth, rightChildWidth))
	}
	return dfs(root, 1, 1)
}
