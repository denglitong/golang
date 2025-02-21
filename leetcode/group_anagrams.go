package leetcode

import (
	"slices"
)

func hashKey(str string) string {
	bytes := []byte(str)
	ints := []int{}

	for _, b := range bytes {
		ints = append(ints, int(b))
	}

	slices.Sort(bytes)

	return string(bytes)
}

func hashKeyV2(str string) string {
	bytes := []byte(str)
	charCount := [26]int{}

	for _, b := range bytes {
		charCount[b-'a']++
	}

	sortedBytes := []byte{}
	for offset, count := range charCount {
		for i := 0; i < count; i++ {
			sortedBytes = append(sortedBytes, byte(offset+'a'))
		}
	}

	return string(sortedBytes)
}

// https://leetcode.cn/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	groups := make(map[string][]string)

	for _, str := range strs {
		key := hashKeyV2(str)
		if _, ok := groups[key]; !ok {
			groups[key] = []string{}
		}
		groups[key] = append(groups[key], str)
	}

	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
