package miscellaneous

import "testing"

type testCaseCountPartitionsEvenSum struct {
	testName string
	input    []int
	result   int
}

func TestCountPartitionsEvenSum(t *testing.T) {
	tests := []testCaseCountPartitionsEvenSum{
		{"TestCase1", []int{10, 10, 3, 7, 6}, 4},
		{"TestCase1", []int{1, 2, 2}, 0},
		{"TestCase1", []int{2, 4, 6, 8}, 3},
	}

	for _, test := range tests {
		results := CountPartitionsEvenSum(test.input)
		if results != test.result {
			t.Errorf("%s: got %d, want %d", test.testName, results, test.result)
		}
	}
}
