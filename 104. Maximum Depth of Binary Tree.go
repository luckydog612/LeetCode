package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	i := 0
	j := 0
	if root == nil {
		return 0
	}
	i = maxDepth(root.Left) + 1
	j = maxDepth(root.Right) + 1
	if i < j {
		i, j = j, i
	}
	return i
}
