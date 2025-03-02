package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node) *Node {
	copyList := &Node{}
	var copyTail, copyPrev *Node = copyList, nil
	listToCopy := make(map[*Node]*Node)

	for ptr := head; ptr != nil; ptr = ptr.Next {
		newNode := &Node{Val: ptr.Val}
		listToCopy[ptr] = newNode

		copyTail.Next = newNode
		copyTail = copyTail.Next

		if copyPrev != nil {
			copyPrev.Next = newNode
		}
		copyPrev = newNode
	}
	for ptr, copyPtr := head, copyList.Next; ptr != nil; ptr, copyPtr = ptr.Next, copyPtr.Next {
		copyPtr.Random = listToCopy[ptr.Random]
	}
	return copyList.Next
}
