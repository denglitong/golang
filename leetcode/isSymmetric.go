package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/symmetric-tree
func isSymmetric(root *TreeNode) bool {
	emptyNode := &TreeNode{Val: 128}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		arr := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			arr = append(arr, queue[i].Val)
			if queue[i].Left == nil {
				queue = append(queue, emptyNode)
			} else {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right == nil {
				queue = append(queue, emptyNode)
			} else {
				queue = append(queue, queue[i].Right)
			}
		}
		ok, hasAtLeastNum := isArrSymmetric(arr)
		// fmt.Println(arr, queue[length:], ok, hasAtLeastNum)
		if !ok {
			return false
		} else if !hasAtLeastNum {
			return ok
		}
		queue = queue[length:]
	}
	return true
}

func isArrSymmetric(arr []int) (bool, bool) {
	if len(arr) == 1 {
		return true, arr[0] != 128
	}
	l, r := 0, len(arr)-1
	hasAtLeastNum := false
	for l < r {
		if arr[l] != 128 || arr[r] != 128 {
			hasAtLeastNum = true
		}
		if arr[l] != arr[r] {
			return false, hasAtLeastNum
		}
		l++
		r--
	}
	return true, hasAtLeastNum
}
