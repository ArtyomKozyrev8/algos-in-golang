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
