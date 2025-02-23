package leetcode

// https://leetcode.cn/problems/min-stack/description
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	// TODO:
	if len(s.min) == 0 {
		s.min = append(s.min, val)
	} else {
		minVal := s.min[len(s.min)-1]
		if val < minVal {
			minVal = val
		}
		s.min = append(s.min, minVal)
	}
	// fmt.Println(s.min)
}

func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.min = s.min[:len(s.min)-1]
	// fmt.Println(s.min)
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.data)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
