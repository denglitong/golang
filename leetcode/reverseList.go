package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	reversed := reverseList(head.Next)
	node.Next = nil
	tail := reversed
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return reversed
}

// 借助一个空节点、前节点，O(N) 复杂度，链表的双指针法
func reverseListV2(head *ListNode) *ListNode {
	var prev, curr, next *ListNode = nil, head, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
