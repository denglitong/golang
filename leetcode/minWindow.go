package leetcode

func windowContain(a [52]int, b [52]int) bool {
	for k := 0; k < 52; k++ {
		if a[k] < b[k] {
			return false
		}
	}
	return true
}

func ch2idx(ch byte) int {
	if 'A' <= ch && ch <= 'Z' {
		return int(ch - 'A')
	}
	return int(26 + ch - 'a')
}

func str2win(s string) [52]int {
	win := [52]int{}
	for _, ch := range s {
		win[ch2idx(byte(ch))]++
	}
	return win
}

func diffWin(s_wins [][52]int, i, j int) [52]int {
	if i < 0 {
		return s_wins[j]
	}
	res := [52]int{}
	for k := 0; k < 52; k++ {
		res[k] = s_wins[j][k] - s_wins[i][k]
	}
	return res
}

// https://leetcode.cn/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {
	s_len, t_len := len(s), len(t)
	if s_len < t_len {
		return ""
	}
	t_win := str2win(t)
	s_wins := make([][52]int, s_len)
	win := [52]int{}
	for i := 0; i < s_len; i++ {
		win[ch2idx(s[i])]++
		s_wins[i] = win
	}

	res := ""
	for i := 0; i <= s_len-t_len; i++ {
		for j := i + t_len - 1; j < s_len; j++ {
			s_win := diffWin(s_wins, i-1, j)
			if windowContain(s_win, t_win) {
				if res == "" || j+1-i < len(res) {
					res = s[i : j+1]
				}
				if j+1-i == t_len {
					return res
				}
				break
			}
		}
	}
	return res
}
