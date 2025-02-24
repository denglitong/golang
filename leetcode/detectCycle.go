package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
			if p1 == p2 {
				break
			}
		}
	}
	if p1 == p2 && p1 != nil {
		p1 = head
		for p1 != p2 {
			p1 = p1.Next
			p2 = p2.Next
		}
		return p1
	}
	return nil
}
