package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p1, p2 *ListNode = head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	var prev *ListNode
	for p2 != nil {
		p2 = p2.Next
		prev = p1
		p1 = p1.Next
	}
	if prev == nil {
		head = p1.Next
	} else {
		prev.Next = p1.Next
	}
	return head
}
