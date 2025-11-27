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

func FindFirstIntersection(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[0] {
			fullOk := true
			for i := 1; i < len(needle); i++ {
				if j+i >= len(haystack) || needle[i] != haystack[j+i] {
					fullOk = false
					break
				}
			}
			if fullOk {
				return j
			}
		}
	}

	return -1
}
