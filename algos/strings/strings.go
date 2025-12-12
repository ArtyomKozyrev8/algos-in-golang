package strings

import (
	"maps"
)

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

func FirstUniqueChar(s string) int {
	values := make(map[int32][2]int)
	for i, val := range s {
		res, exist := values[val]
		if exist == false {
			// we store two values, first value is Index when letter first met
			// second is number of such letters
			values[val] = [2]int{i, 1}
		} else {
			values[val] = [2]int{res[0], res[1] + 1}
		}
	}

	minIndex := -1
	// map does not keep keys order
	for _, v := range values {
		if v[1] == 1 {
			if minIndex == -1 {
				minIndex = v[0]
			} else {
				if v[0] < minIndex {
					minIndex = v[0]
				}
			}
		}
	}

	return minIndex
}

func RepeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	windowSize := 1
	maxWindowSize := (len(s) - len(s)%2) / 2

	for windowSize <= maxWindowSize {
		subString := s[0:windowSize]
		allFound := true
		for i := 0; i < len(s); {
			if i+windowSize > len(s) {
				allFound = false
				break
			}

			temp := s[i : i+windowSize]
			if subString != temp {
				allFound = false
				break
			} else {
				i = i + windowSize
			}
		}

		if allFound {
			return true
		} else {
			windowSize += 1
		}
	}

	return false
}

func FindAnagramsWayOne(s string, p string) []int {
	var results []int

	var pNum int = 0                  // sum of chars
	pComposition := make(map[int]int) // composition of letters un p

	for i := 0; i < len(p); i++ {
		pNum += int(p[i])
		if _, exists := pComposition[int(p[i])]; !exists {
			pComposition[int(p[i])] = 1
		} else {
			pComposition[int(p[i])] += 1
		}
	}

	isFirstTime := true
	var curWindowNum int = 0

	for i := 0; i < len(s); i++ {
		if i+len(p)-1 >= len(s) {
			break
		}

		if isFirstTime {
			isFirstTime = false
			for j := 0; j < len(s[i:i+len(p)]); j++ {
				curWindowNum += int(s[i : i+len(p)][j])
			}
		} else {
			charVal := int(s[i+len(p)-1])
			curWindowNum = curWindowNum + charVal
		}

		if curWindowNum == pNum {
			allSame := true
			localMapComposition := make(map[int]int)
			maps.Copy(localMapComposition, pComposition)
			for j := 0; j < len(s[i:i+len(p)]); j++ {
				val, exists := localMapComposition[int(s[i : i+len(p)][j])]
				localMapComposition[int(s[i : i+len(p)][j])] = val - 1
				if val-1 < 0 {
					allSame = false
					break
				}

				if !exists {
					allSame = false
					break
				}
			}
			if allSame {
				results = append(results, i)
			}
		}
		minusCharVal := int(s[i])
		curWindowNum -= minusCharVal
	}

	return results
}

func FindAnagramsWayTwo(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	pNum := 0
	pComposition := make(map[int]int)

	for i := 0; i < len(p); i++ {
		pNum += int(p[i])
		if _, exists := pComposition[int(p[i])]; !exists {
			pComposition[int(p[i])] = 1
		} else {
			pComposition[int(p[i])] += 1
		}
	}

	curNum := 0
	curComposition := make(map[int]int)
	var results []int

	for i := 0; i < len(s); i++ {
		if i+len(p) > len(s) {
			break
		}

		window := s[i : i+len(p)]
		if i == 0 {
			for j := 0; j < len(window); j++ {
				curNum += int(s[j])
				if _, exists := curComposition[int(window[j])]; !exists {
					curComposition[int(window[j])] = 1
				} else {
					curComposition[int(window[j])] += 1
				}
			}
		} else {
			toRemove := int(s[i-1])
			toAdd := int(s[i+len(p)-1])

			if _, exists := curComposition[toAdd]; !exists {
				curComposition[toAdd] = 1
			} else {
				curComposition[toAdd] += 1
			}
			curNum += toAdd

			if _, exists := curComposition[toRemove]; !exists {
				// it can't be so ...
			} else {
				curComposition[toRemove] -= 1
				if curComposition[toRemove] <= 0 {
					delete(curComposition, toRemove)
				}
			}
			curNum -= toRemove
		}

		if curNum == pNum {
			if len(curComposition) == len(pComposition) {
				if maps.Equal(curComposition, pComposition) {
					results = append(results, i)
				}

			}
		}
	}

	return results
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	reversedS := reverseString(s)

	windowSize := len(s)
	for {
		j := len(s) - windowSize
		for i := 0; i < len(s)-windowSize+1; i++ {
			curS := s[i : i+windowSize]
			curR := reversedS[j : j+windowSize]
			if curS == curR {
				return curS
			}
			j--
		}

		windowSize -= 1
	}
}

func IsValidParenthesesOrderString(s string) bool {
	var storage []rune

	charMap := map[rune]int{
		'(': 1, '[': 2, '{': 3, ')': -1, ']': -2, '}': -3,
	}

	for _, char := range s {
		val, exists := charMap[char]
		if !exists {
			return false // unexpected value
		}
		if val > 0 {
			storage = append(storage, char)
		} else {
			if len(storage) == 0 {
				return false
			}
			lastChar := storage[len(storage)-1]
			if charMap[lastChar] != val*-1 {
				return false
			}
			storage = storage[:len(storage)-1]
		}
	}

	if len(storage) != 0 {
		return false
	}
	return true
}

func CheckSimpleParentheses(s string) bool {
	var chars []rune
	for _, char := range s {
		if char == '(' {
			chars = append(chars, '(')
		} else {
			if len(chars) == 0 {
				return false
			}
			chars = chars[:len(chars)-1]
		}
	}

	return len(chars) == 0
}

func WordBreak(s string, wordDict []string) bool {
	visitedIndexes := make(map[int]bool, len(wordDict)) // allows not to check the same strings again
	startIndexes := []int{0}                            // we do not need to store the whole string, only index of first letter is enough
	for len(startIndexes) > 0 {
		var newStartIndexes []int
		for _, index := range startIndexes {
			for _, word := range wordDict {
				if len(s[index:]) >= len(word) && word == s[index:index+len(word)] {
					newIndex := index + len(word)
					if newIndex == len(s) {
						return true
					}
					if !visitedIndexes[newIndex] {
						visitedIndexes[newIndex] = true
						newStartIndexes = append(newStartIndexes, newIndex)
					}
				}
			}
		}
		if len(newStartIndexes) == 0 {
			return false
		}
		startIndexes = newStartIndexes
	}
	return true
}

func LengthOfLastWord(s string) int {
	rightIndex := -1
	leftIndex := -1
	space := " "[0]
	for i := len(s) - 1; i >= 0; i-- {
		if rightIndex > -1 && leftIndex > -1 {
			break
		}

		if s[i] != space {
			if rightIndex == -1 {
				rightIndex = i
			}
		} else {
			if leftIndex == -1 {
				if rightIndex != -1 {
					leftIndex = i
				}
			}
		}
	}

	if rightIndex == -1 {
		return 0
	}

	if leftIndex == -1 {
		return len(s[:rightIndex+1])
	}

	return len(s[leftIndex:rightIndex])
}
