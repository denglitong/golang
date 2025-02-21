package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		inputStrs   []string
		expectedArr [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
		{
			[]string{"a"},
			[][]string{
				{"a"},
			},
		},
	}

	for _, testCase := range testCases {
		actualVal := groupAnagrams(testCase.inputStrs)
		assert.Equal(t, testCase.expectedArr, actualVal)
	}
}
