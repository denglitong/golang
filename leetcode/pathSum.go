package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/path-sum-iii/description/
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	preSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		res += preSum[curr-targetSum]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return res
}
