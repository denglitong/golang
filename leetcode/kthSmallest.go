package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	arr := inorder(root)
	fmt.Println(arr)
	return arr[k-1]
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorder(root.Left)
	left = append(left, root.Val)
	right := inorder(root.Right)
	return append(left, right...)
}
