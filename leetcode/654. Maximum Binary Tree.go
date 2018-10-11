package LeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	t := &TreeNode{}
	if len(nums) == 0 {
		return nil
	}
	max := 0
	for _, i := range nums {
		if i > max {
			i, max = max, i
		}
	}
	k := 0
	for j := range nums {
		if nums[j] == max {
			k = j
		}
	}
	left := nums[:k]
	right := nums[k+1:]
	t.Val = max
	t.Left = constructMaximumBinaryTree(left)
	t.Right = constructMaximumBinaryTree(right)
	return t
}
