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

type testClimbStairsRecursive struct {
	testName string
	number   int
	result   int
}

func TestClimbStairsRecursive(t *testing.T) {
	tests := []testClimbStairsRecursive{
		{"TestCase1", 3, 3},
		{"TestCase2", 1, 1},
		{"TestCase3", 2, 2},
		{"TestCase4", 4, 5},
	}

	for _, test := range tests {
		results := ClimbStairsRecursive(test.number)
		if results != test.result {
			t.Errorf("%s: got %d, want %d", test.testName, results, test.result)
		}
	}
}
