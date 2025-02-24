package leetcode

// https://leetcode.cn/problems/binary-tree-level-order-traversal
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		arr := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			arr = append(arr, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, arr)
		queue = queue[length:]
	}
	return res
}
