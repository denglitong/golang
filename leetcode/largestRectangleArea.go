package leetcode

type IncQueue struct {
	// [2]int{val, index]
	data [][2]int
}

func NewIncQueue(init [2]int) *IncQueue {
	return &IncQueue{
		data: [][2]int{init},
	}
}

func (q *IncQueue) Push(item [2]int) int {
	for len(q.data) > 0 && q.Top()[0] >= item[0] {
		q.Pop()
	}
	leftIdx := q.data[len(q.data)-1][1]
	q.data = append(q.data, item)
	return leftIdx
}

func (q *IncQueue) Top() [2]int {
	if len(q.data) == 0 {
		return [2]int{-1, -1}
	}
	return q.data[len(q.data)-1]
}

func (q *IncQueue) Pop() {
	if len(q.data) > 0 {
		q.data = q.data[:len(q.data)-1]
	}
}

// https://leetcode.cn/problems/largest-rectangle-in-histogram
func largestRectangleArea(heights []int) int {
	res, n := heights[0], len(heights)
	leftQueue := NewIncQueue([2]int{-1, -1}) // [2]{val, idx}
	leftIdx := make([]int, n+1)
	rightQueue := NewIncQueue([2]int{-1, n}) // [2]{val, idx}
	rightIdx := make([]int, n+1)

	for i := 0; i < n; i++ {
		idx := leftQueue.Push([2]int{heights[i], i})
		leftIdx[i] = idx
	}
	for i := n - 1; i >= 0; i-- {
		idx := rightQueue.Push([2]int{heights[i], i})
		rightIdx[i] = idx
	}

	for i := 0; i < n; i++ {
		area := (rightIdx[i] - leftIdx[i] - 1) * heights[i]
		if area > res {
			res = area
		}
	}
	//fmt.Println(heights)
	//fmt.Println(leftIdx)
	//fmt.Println(rightIdx)
	return res
}
