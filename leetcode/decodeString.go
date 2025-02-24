package leetcode

import (
	"strings"
)

func isDigit(ch byte) (int, bool) {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0'), true
	}
	return 0, false
}

// https://leetcode.cn/problems/decode-string/description/
func decodeString(s string) string {
	// cnt1,indexOfLeft,...,indexOfRight, cnt2,indexOfLeft,...,indexOfRight
	tokenArr := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			leftFlag := tokenArr[len(tokenArr)-1]
			digit := tokenArr[len(tokenArr)-2]
			digitLeftIdx := tokenArr[len(tokenArr)-3]
			tokenArr = tokenArr[:len(tokenArr)-3]

			repeats := []string{}
			for j := 0; j < digit; j++ {
				repeats = append(repeats, s[leftFlag+1:i])
			}
			decode := strings.Join(repeats, "")
			s = strings.Join([]string{
				s[0:digitLeftIdx], decode, s[i+1:],
			}, "")
			//fmt.Println(digitLeftIdx, digit, leftFlag, i, decode, s)
			i = digitLeftIdx + len(decode) - 1
			continue
		}

		d, ok := isDigit(s[i])
		if !ok {
			continue
		}

		tokenArr = append(tokenArr, i)
		cnt := 0
		for i < len(s) && ok {
			cnt = 10*cnt + d
			//fmt.Println("f1", cnt)
			i++
			d, ok = isDigit(s[i])
		}
		tokenArr = append(tokenArr, cnt)
		tokenArr = append(tokenArr, i)
	}

	return s
}
