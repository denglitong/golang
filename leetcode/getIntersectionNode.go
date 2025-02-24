package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/intersection-of-two-linked-lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptr1, ptr2 := headA, headB
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
		if ptr1 == ptr2 && ptr1 != nil {
			return ptr1
		} else if ptr1 == ptr2 && ptr1 == nil {
			return nil
		} else if ptr1 == nil {
			ptr1 = headB
		} else if ptr2 == nil {
			ptr2 = headA
		}
	}
	return ptr1
}
