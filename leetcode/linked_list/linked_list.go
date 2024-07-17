package linked_list

// Definition for singly-linked inputList.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
func getDecimalValue(head *ListNode) int {
	res := 0
	for p := head; p != nil; p = p.Next {
		res = res*2 + p.Val
	}
	return res
}
