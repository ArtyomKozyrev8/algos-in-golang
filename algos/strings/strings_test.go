package strings

import "testing"

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
