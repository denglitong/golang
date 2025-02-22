package leetcode

type MonotonicDecreaseQueue struct {
	data []int
}

func NewMonotonicDecreaseQueue() *MonotonicDecreaseQueue {
	return &MonotonicDecreaseQueue{data: []int{}}
}

func (s *MonotonicDecreaseQueue) Push(v int) {
	for len(s.data) > 0 && s.data[len(s.data)-1] < v {
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, v)
}

func (s *MonotonicDecreaseQueue) Pop(val int) {
	if val == s.data[0] {
		s.data = s.data[1:]
	}
}

func (s *MonotonicDecreaseQueue) Max() int {
	return s.data[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	mdq := NewMonotonicDecreaseQueue()
	for i := 0; i < k; i++ {
		mdq.Push(nums[i])
	}
	res := []int{mdq.Max()}
	n := len(nums)
	for j := k; j < n; j++ {
		mdq.Push(nums[j])
		mdq.Pop(nums[j-k])
		res = append(res, mdq.Max())
	}
	return res
}

type MonotonicIncreaseQueue struct {
	data []int
}

func NewMonotonicIncreaseQueue() *MonotonicIncreaseQueue {
	return &MonotonicIncreaseQueue{data: []int{}}
}

func (s *MonotonicIncreaseQueue) Push(v int) {
	for len(s.data) > 0 && s.data[len(s.data)-1] > v {
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, v)
}

func (s *MonotonicIncreaseQueue) Pop(val int) {
	if val == s.data[0] {
		s.data = s.data[1:]
	}
}

func (s *MonotonicIncreaseQueue) Min() int {
	return s.data[0]
}

func maxSlidingWindowV2(nums []int, k int) []int {
	miq := NewMonotonicIncreaseQueue()
	for i := 0; i < k; i++ {
		miq.Push(nums[i])
	}
	res := []int{miq.Min()}
	n := len(nums)
	for j := k; j < n; j++ {
		miq.Push(nums[j])
		miq.Pop(nums[j-k])
		res = append(res, miq.Min())
	}
	return res
}
