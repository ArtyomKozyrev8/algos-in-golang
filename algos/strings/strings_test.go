package strings

import (
	"slices"
	"testing"
)

func TestLengthOfLongestSubstringWayOne(t *testing.T) {
	var tests = struct {
		input []string
		want  []int
	}{
		[]string{"abcabcbb", "bbbbb", "pwwkew", ""},
		[]int{3, 1, 3, 0},
	}

	for i := 0; i < len(tests.want); i++ {
		res := LengthOfLongestSubstringWayOne(tests.input[i])
		if res != tests.want[i] {
			t.Errorf("%d != %d", res, tests.want[i])
		}
	}
}

type testCaseFindFirstIntersection struct {
	caseName string
	haystack string
	needle   string
	result   int
}

func TestFindFirstIntersection(t *testing.T) {
	var tests = []testCaseFindFirstIntersection{
		{"TestCase-1", "sadbutsad", "sad", 0},
		{"TestCase-2", "leetcode", "leeto", -1},
		{"TestCase-3", "mississippi", "issipi", -1},
		{"TestCase-3", "issipi", "mississippi", -1},
	}

	for _, testCase := range tests {
		res := FindFirstIntersection(testCase.haystack, testCase.needle)
		if res != testCase.result {
			t.Errorf("%s %d != %d", testCase.caseName, res, testCase.result)
		}
	}
}

type testCaseIsSubsequence struct {
	caseName      string
	s             string
	t             string
	isSubsequence bool
}

func TestIsSubsequence(t *testing.T) {
	var tests = []testCaseIsSubsequence{
		{"TestCase-1", "abc", "ahbgdc", true},
		{"TestCase-2", "axc", "ahbgdc", false},
		{"TestCase-3", "aaaaaa", "bbaaaa", false},
		{"TestCase-3", "aa", "a", false},
	}

	for _, testCase := range tests {
		res := IsSubsequence(testCase.s, testCase.t)
		if res != testCase.isSubsequence {
			t.Errorf("%s %v != %v", testCase.caseName, testCase.isSubsequence, res)
		}
	}
}

type testCaseIsSubstring struct {
	caseName string
	s        string
	result   int
}

func TestFirstUniqueChar(t *testing.T) {
	tests := []testCaseIsSubstring{
		{"TestCase-1", "abc", 0},
		{"TestCase-2", "aabb", -1},
		{"TestCase-2", "helloworld", 0},
	}

	for _, testCase := range tests {
		res := FirstUniqueChar(testCase.s)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type testRepeatedSubstringPattern struct {
	caseName string
	s        string
	result   bool
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []testRepeatedSubstringPattern{
		{"TestCase-1", "ac", false},
		{"TestCase-2", "a", true},
		{"TestCase-3", "", false},
		{"TestCase-4", "abab", true},
		{"TestCase-5", "abcabcabc", true},
		{"TestCase-5", "abcab", false},
	}
	for _, testCase := range tests {
		res := RepeatedSubstringPattern(testCase.s)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type TestFindAnagrams struct {
	caseName string
	s        string
	p        string
	result   []int
}

func TestFindAnagramsWayOne(t *testing.T) {
	tests := []TestFindAnagrams{
		{"TestCase-1", "abc", "abc", []int{0}},
		{"TestCase-2", "cbaebabacd", "abc", []int{0, 6}},
		{"TestCase-3", "abab", "ab", []int{0, 1, 2}},
		{"TestCase-4", "af", "eb", []int{}},
		{"TestCase-5", "af", "ebc", []int{}},
		{"TestCase-6", "abcde", "a", []int{0}},
	}

	for _, testCase := range tests {
		res := FindAnagramsWayOne(testCase.s, testCase.p)
		if !slices.Equal(res, testCase.result) {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

func TestFindAnagramsWayTwo(t *testing.T) {
	tests := []TestFindAnagrams{
		{"TestCase-1", "abc", "abc", []int{0}},
		{"TestCase-2", "cbaebabacd", "abc", []int{0, 6}},
		{"TestCase-3", "abab", "ab", []int{0, 1, 2}},
		{"TestCase-4", "af", "eb", []int{}},
		{"TestCase-5", "af", "ebc", []int{}},
		{"TestCase-6", "abcde", "a", []int{0}},
	}

	for _, testCase := range tests {
		res := FindAnagramsWayTwo(testCase.s, testCase.p)
		if !slices.Equal(res, testCase.result) {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type testLongestPalindromeTwo struct {
	caseName string
	s        string
	result   string
}

func TestLongestPalindrome(t *testing.T) {
	tests := []testLongestPalindromeTwo{
		{"TestCase-1", "babad", "bab"},
		{"TestCase-2", "cbbd", "bb"},
		{"TestCase-3", "abb", "bb"},
		{"TestCase-4", "a", "a"},
		{"TestCase-5", "", ""},
		{"TestCase-6", "abcdef", "a"},
	}

	for _, testCase := range tests {
		res := LongestPalindrome(testCase.s)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type testIsValidParenthesesOrderString struct {
	caseName string
	s        string
	result   bool
}

func TestIsValidParenthesesOrderString(t *testing.T) {
	tests := []testIsValidParenthesesOrderString{
		{"TestCase-1", "()", true},
		{"TestCase-2", "()[]{}", true},
		{"TestCase-3", "(]", false},
		{"TestCase-4", "([])", true},
		{"TestCase-5", "([)]", false},
		{"TestCase-6", "]", false},
		{"TestCase-7", "(", false},
		{"TestCase-7", "(-)", false},
	}

	for _, testCase := range tests {
		res := IsValidParenthesesOrderString(testCase.s)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type testCheckSimpleParentheses struct {
	caseName string
	s        string
	result   bool
}

func TestCheckSimpleParentheses(t *testing.T) {
	tests := []testCheckSimpleParentheses{
		{"TestCase-1", "()", true},
		{"TestCase-2", "(())", true},
		{"TestCase-3", "(", false},
		{"TestCase-4", ")", false},
		{"TestCase-5", "())", false},
		{"TestCase-6", "(()", false},
	}

	for _, testCase := range tests {
		res := CheckSimpleParentheses(testCase.s)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}

type testWordBreak struct {
	caseName string
	s        string
	wordDict []string
	result   bool
}

func TestWordBreak(t *testing.T) {
	tests := []testWordBreak{
		{"TestCase-1", "abc", []string{"abc"}, true},
		{"TestCase-2", "abcabc", []string{"abc"}, true},
		{"TestCase-3", "abcabc", []string{"ab"}, false},
		{"TestCase-4", "abcabc", []string{"ab"}, false},
	}
	for _, testCase := range tests {
		res := WordBreak(testCase.s, testCase.wordDict)
		if res != testCase.result {
			t.Errorf("%s %v != %v", testCase.caseName, res, testCase.result)
		}
	}
}
