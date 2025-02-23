package leetcode

func check(w1 map[uint8]int, w2 map[uint8]int) bool {
	for ch2, cnt2 := range w2 {
		cnt1, ok := w1[ch2]
		if !ok || cnt1 < cnt2 {
			return false
		}
	}
	return true
}

func minWindowV2(s string, t string) string {
	res := ""
	sLen, tLen := len(s), len(t)
	if sLen < tLen {
		return res
	}

	winS, winT := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < tLen; i++ {
		winS[s[i]]++
		winT[t[i]]++
	}
	if check(winS, winT) {
		return s[0:tLen]
	}

	l, r := 0, tLen
	for r < sLen {
		winS[s[r]]++
		//fmt.Println("f0", winS)
		for l <= r && check(winS, winT) {
			if len(res) == 0 || r-l+1 < len(res) {
				//fmt.Println("f1", s[l:r+1])
				res = s[l : r+1]
			}
			winS[s[l]]--
			l++
		}
		r++
	}
	return res
}
