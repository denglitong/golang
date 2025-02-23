package leetcode

// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	//digits := make(map[int][]int)

	//for i, num := range nums {
	//	if _, ok := digits[num]; !ok {
	//		digits[num] = []int{i}
	//	} else {
	//		digits[num] = append(digits[num], i)
	//	}
	//}
	//
	//for i, num := range nums {
	//	idxArr, ok := digits[target-num]
	//	if !ok {
	//		continue
	//	}
	//	//fmt.Println(nums, num, idxArr)
	//	if len(idxArr) == 1 && idxArr[0] != i {
	//		return []int{i, idxArr[0]}
	//	}
	//	if len(idxArr) == 2 && target-num == num {
	//		return idxArr
	//	}
	//}

	hashTable := map[int]int{}

	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			return []int{i, j}
		}
		hashTable[num] = i
	}

	return nil
}
