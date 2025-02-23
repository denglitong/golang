package leetcode

func hash(bytes []byte) string {
	letters := [26]int{}
	for _, b := range bytes {
		letters[int(b-'a')]++
	}
	sorted_bytes := []byte{}
	for i := 0; i < 26; i++ {
		for letters[i] > 0 {
			sorted_bytes = append(sorted_bytes, byte('a'+i))
			letters[i]--
		}
	}
	return string(sorted_bytes)
}

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
func findAnagrams(s string, p string) []int {
	s_len, p_len := len(s), len(p)
	if p_len > s_len {
		return nil
	}

	s_window, p_window := [26]byte{}, [26]byte{}
	for i, ch := range []byte(p) {
		p_window[ch-'a']++
		s_window[s[i]-'a']++
	}

	res := []int{}
	if s_window == p_window {
		res = append(res, 0)
	}

	for i := 1; i <= s_len-p_len; i++ {
		s_window[s[i-1]-'a']--
		s_window[s[i+p_len-1]-'a']++
		if s_window == p_window {
			res = append(res, i)
		}
	}

	return res
}
