package leetcode

import (
	"fmt"
	"testing"
)

func TestMinWindowV2(t *testing.T) {
	fmt.Println(minWindowV2("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindowV2("abc", "abc"))
	fmt.Println(minWindowV2("a", "a"))
	fmt.Println(minWindowV2("ab", "b"))
}
