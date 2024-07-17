package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	testCases := []struct {
		inputList   *ListNode
		expectedVal int
	}{
		{linkedListBuilder([]int{1, 0, 1}), 5},
		{linkedListBuilder([]int{0}), 0},
		{linkedListBuilder([]int{1}), 1},
		{linkedListBuilder([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), 18880},
		{linkedListBuilder([]int{0, 0}), 0},
	}

	for _, testCase := range testCases {
		actualVal := getDecimalValue(testCase.inputList)
		assert.Equal(t, testCase.expectedVal, actualVal)
	}
}

func linkedListBuilder(vals []int) *ListNode {
	list := &ListNode{}
	head := list
	for _, val := range vals {
		list.Next = &ListNode{Val: val}
		list = list.Next
	}
	return head.Next
}
