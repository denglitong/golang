package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lead := &ListNode{Next: head}
	prev, curr, next := lead, head, head.Next
	var tmp *ListNode
	for curr != nil && next != nil {
		prev.Next = next
		tmp = next.Next
		next.Next = curr
		curr.Next = tmp

		prev = curr
		curr = tmp
		if curr != nil {
			next = curr.Next
		}
	}
	return lead.Next
}
