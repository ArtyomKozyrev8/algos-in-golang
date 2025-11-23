package strings

// LengthOfLongestSubstringWayOne  - we do not want to use rune, we have only english letters in string
func LengthOfLongestSubstringWayOne(s string) int {
	longestUniqueCharsSubstringsLen := 0
	LenS := len(s)

	for i := 0; i < LenS; {
		if LenS-i <= longestUniqueCharsSubstringsLen {
			break
		}
		currentLongestChars := make(map[uint8]int)
		lenCurrentLongestChars := 0

		for j := i; j < LenS; j++ {
			pos, exist := currentLongestChars[s[j]]
			if exist {
				i = pos + 1
				break
			} else {
				currentLongestChars[s[j]] = j
				lenCurrentLongestChars += 1
				if lenCurrentLongestChars > longestUniqueCharsSubstringsLen {
					longestUniqueCharsSubstringsLen = lenCurrentLongestChars
				}
				i++
			}
			if LenS-j+lenCurrentLongestChars <= longestUniqueCharsSubstringsLen {
				break
			}
		}
	}

	return longestUniqueCharsSubstringsLen
}
