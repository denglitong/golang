package leetcode

func lengthOfLongestSubstring(s string) int {
	l, r, n, curr, maxLen := 0, 0, len(s), 0, 0
	window := make(map[byte]int)
	for r < n {
		if idx, ok := window[s[r]]; ok {
			for l <= idx {
				delete(window, s[l])
				curr--
				l++
			}
		}
		window[s[r]] = r
		curr++
		if curr > maxLen {
			maxLen = curr
		}
		r++
	}
	return maxLen
}
