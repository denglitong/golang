package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, Metadata{})
}

func isValidBSTHelper(root *TreeNode, meta Metadata) bool {
	if root == nil {
		return true
	}
	if (meta.isMinSet && root.Val <= meta.min) ||
		(meta.isMaxSet && root.Val >= meta.max) {
		return false
	}
	if (root.Left != nil && root.Left.Val >= root.Val) ||
		(root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}

	leftMeta, rightMeta := meta, meta
	leftMeta.max = root.Val
	leftMeta.isMaxSet = true
	rightMeta.min = root.Val
	rightMeta.isMinSet = true

	return isValidBSTHelper(root.Left, leftMeta) && isValidBSTHelper(root.Right, rightMeta)
}

type Metadata struct {
	min, max           int
	isMinSet, isMaxSet bool
}
