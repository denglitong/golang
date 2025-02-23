package leetcode

import "fmt"

type LCS struct {
	sequences  [][]int
	rangeIndex map[int]int
}

func NewLCS() *LCS {
	return &LCS{
		sequences:  make([][]int, 0),
		rangeIndex: make(map[int]int),
	}
}

func (s *LCS) Push(val int) {
	// 1. if sequences is empty, push and return
	// 2. if sequences is not empty,
	// 	find if val is within some sequence seq[i],
	//		if it's not equal to seq[i][1]+1, do nothing and return
	// 		if it's equal to seq[i][1]+1, update seq[i][1] = seq[i][1]+1
	// 			if seq[i][1] is not equal to seq[i+1][0], do nothing and return
	// 			if seq[i][1] is equal to seq[i+1][0],
	//				update seq[i][1] = seq[i+1][1], delete the row seq[i+1], and return
	//  if val is not within at any sequence, insert a new sequence in the position:
	//		seq[i][1] < val < seq[i+1][0]

	if len(s.sequences) == 0 {
		s.sequences = append(s.sequences, []int{val, val})
		s.rangeIndex[val] = 0
		// fmt.Println(val, s.sequences, "f1")
		return
	}

	if val < s.sequences[0][0]-1 {
		s.sequences = append([][]int{{val, val}}, s.sequences...)
		// fmt.Println(val, s.sequences, "f2")
		return
	} else if val == s.sequences[0][0]-1 {
		s.sequences[0][0] = val
		// fmt.Println(val, s.sequences, "f3")
		return
	} else {
		seqLen := len(s.sequences)

		for i := 0; i < seqLen; i++ {
			if s.sequences[i][0] <= val && val <= s.sequences[i][1] {
				// fmt.Println(val, s.sequences, "f4")
				return
			}

			if i == seqLen-1 {
				// fmt.Println(s.sequences, "f6-0")
				if val == s.sequences[i][1]+1 {
					s.sequences[i][1] = val
					// fmt.Println(val, s.sequences, "f5")
					return
				}
				// fmt.Println(s.sequences, "f6-1")
				s.sequences = append(s.sequences, []int{val, val})
				// fmt.Println(val, s.sequences, "f6-2")
				return
			}

			if s.sequences[i][1]+1 < val && val < s.sequences[i+1][0]-1 {
				right := append([][]int{}, s.sequences[i+1:]...)
				left := append(s.sequences[:i+1], []int{val, val})
				s.sequences = append(left, right...)
				// fmt.Println(val, s.sequences, "f7")
				return
			}

			if s.sequences[i][1]+1 == val && val < s.sequences[i+1][0]-1 {
				s.sequences[i][1] = val
				// fmt.Println(val, s.sequences, "f7.1")
				return
			}
			if s.sequences[i][1]+1 < val && s.sequences[i+1][0]-1 == val {
				s.sequences[i+1][0] = val
				// fmt.Println(val, s.sequences, "f7.2")
				return
			}
			if s.sequences[i][1]+1 == val && val == s.sequences[i+1][0]-1 {
				s.sequences[i][1] = s.sequences[i+1][1]
				s.sequences = append(s.sequences[:i+1], s.sequences[i+2:]...)
				// fmt.Println(val, s.sequences, "f8")
				return
			}
		}
	}
}

func (s *LCS) Sequences() [][]int {
	return s.sequences
}

func (s *LCS) MaxLCSLen() int {
	maxLen := 0

	for _, seq := range s.sequences {
		l := seq[1] - seq[0] + 1
		if l > maxLen {
			maxLen = l
		}
	}

	return maxLen
}

// https://leetcode.cn/problems/longest-consecutive-sequence/description
func longestConsecutive(nums []int) int {
	//lcs := NewLCS()
	//for _, n := range nums {
	//	lcs.Push(n)
	//}
	//fmt.Println(lcs.Sequences())
	//fmt.Println(lcs.MaxLCSLen())
	//return lcs.MaxLCSLen()

	//if len(nums) == 0 {
	//	return 0
	//}
	//
	//digits := map[int]bool{}
	//for _, n := range nums {
	//	digits[n] = true
	//}
	//
	//keys := []int{}
	//for k := range digits {
	//	keys = append(keys, k)
	//}
	//
	//slices.Sort(keys)
	//
	//maxLen := 1
	//for i := 0; i < len(keys); i++ {
	//	length := 1
	//	for j := i + 1; j < len(keys) && keys[i]+1 == keys[j]; j++ {
	//		length++
	//		i = j
	//	}
	//	if length > maxLen {
	//		maxLen = length
	//	}
	//}

	maxLen := 0

	digits := map[int]bool{}
	for _, n := range nums {
		digits[n] = true
	}

	for num := range digits {
		if !digits[num-1] {
			length := 1
			currentNum := num
			for digits[currentNum+1] {
				currentNum++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}

	//fmt.Println(keys, digits)
	fmt.Println(maxLen, nums)
	return maxLen
}
