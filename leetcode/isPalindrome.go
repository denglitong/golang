package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	n := 0
	for ptr := head; ptr != nil; ptr = ptr.Next {
		n++
	}
	middle := n / 2
	var prev, curr, next *ListNode = nil, head, nil
	for i := 0; i < middle; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	if n%2 > 0 {
		curr = curr.Next
	}

	for prev != nil && curr != nil {
		if prev.Val != curr.Val {
			return false
		}
		prev = prev.Next
		curr = curr.Next
	}
	return prev == nil && curr == nil
}
