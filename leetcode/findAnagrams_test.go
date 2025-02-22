package leetcode

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
