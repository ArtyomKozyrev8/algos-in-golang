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

// IsSubsequence t is superset of s, but without any letter's position change
func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	data := make(map[uint8][]int, len(t))

	for i := 0; i < len(t); i++ {
		data[t[i]] = append(data[t[i]], i)
	}

	prevFoundPosition := 0
	for i := 0; i < len(s); i++ {
		vals, found := data[s[i]]
		if found {
			positionFound := false
			for _, v := range vals {
				if prevFoundPosition == 0 && v == 0 {
					positionFound = true
					break
				}

				if v > prevFoundPosition {
					prevFoundPosition = v
					positionFound = true
					break
				}
			}
			if positionFound == false {
				return false
			}

		} else {
			return false
		}
	}

	return true
}
