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

// dfs tree
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(height(root.Left), height(root.Right)) + 1
}
