package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	ptr1, ptr2, tail := l1, l2, sum
	increase := 0
	for ptr1 != nil && ptr2 != nil {
		num := ptr1.Val + ptr2.Val + increase
		val := num % 10
		increase = num / 10
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	for ptr1 != nil {
		num := ptr1.Val + increase
		val := num % 10
		increase = num / 10
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
		ptr1 = ptr1.Next
	}
	for ptr2 != nil {
		num := ptr2.Val + increase
		val := num % 10
		increase = num / 10
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
		ptr2 = ptr2.Next
	}
	if increase == 1 {
		tail.Next = &ListNode{Val: 1}
		tail = tail.Next
	}

	sum = sum.Next
	printList(sum)
	return sum
}

func printList(list *ListNode) {
	for ptr := list; ptr != nil; ptr = ptr.Next {
		fmt.Print(ptr.Val)
	}
	fmt.Println()
}
