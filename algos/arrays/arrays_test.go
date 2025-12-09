package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestRemoveDuplicatesFromSortedArrayInPlace(t *testing.T) {
	vars := struct {
		inputs          [][]int
		uniqueValuesNum []int
		resultSlice     [][]int
	}{
		inputs: [][]int{
			{},
			{1},
			{1, 1},
			{1, 1, 1},
			{1, 1, 2, 2, 3, 4, 4, 4, 5},
			{1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5},
			{1, 1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 5, 5},
			{0, 0, 0, 0, 0},
		},
		uniqueValuesNum: []int{0, 1, 1, 1, 5, 5, 5, 1},
		resultSlice: [][]int{
			{}, {1}, {1}, {1}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {0},
		},
	}

	for i := 0; i < len(vars.inputs); i++ {
		uniqueValNam, ResSlice := RemoveDuplicatesFromSortedArrayInPlace(vars.inputs[i])
		if uniqueValNam != vars.uniqueValuesNum[i] {
			t.Errorf("%v != %v", uniqueValNam, vars.uniqueValuesNum[i])
		}
		if !reflect.DeepEqual(ResSlice, vars.resultSlice[i]) {
			t.Errorf("%v != %v", ResSlice, vars.resultSlice[i])
		}
	}
}

func TestRemoveDuplicatesFromSortedArrayInPlaceAttemptTwo(t *testing.T) {
	vars := struct {
		inputs          [][]int
		uniqueValuesNum []int
		resultSlice     [][]int
	}{
		inputs: [][]int{
			{},
			{1},
			{1, 1},
			{1, 1, 1},
			{1, 1, 2, 2, 3, 4, 4, 4, 5},
			{1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5},
			{1, 1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 5, 5},
			{0, 0, 0, 0, 0},
		},
		uniqueValuesNum: []int{0, 1, 1, 1, 5, 5, 5, 1},
		resultSlice: [][]int{
			{}, {1}, {1}, {1}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {0},
		},
	}

	for i := 0; i < len(vars.inputs); i++ {
		uniqueValNam, ResSlice := RemoveDuplicatesFromSortedArrayInPlaceAttemptTwo(vars.inputs[i])
		if uniqueValNam != vars.uniqueValuesNum[i] {
			t.Errorf("%v != %v", uniqueValNam, vars.uniqueValuesNum[i])
		}
		if !reflect.DeepEqual(ResSlice, vars.resultSlice[i]) {
			t.Errorf("%v != %v", ResSlice, vars.resultSlice[i])
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	vars := struct {
		wordsGroup [][]string
		prefixes   []string
	}{
		wordsGroup: [][]string{
			{"abc", "a", "b"}, {"abc", "b", "b"}, {"abc", "abc", "abc"}, {"", "abcd"}, {"abcd", ""}, {}, {"a"},
		},
		prefixes: []string{"", "", "abc", "", "", "", "a"},
	}

	for i := 0; i < len(vars.wordsGroup); i++ {
		prefix := LongestCommonPrefix(vars.wordsGroup[i])
		if prefix != vars.prefixes[i] {
			t.Errorf("%v != %v", prefix, vars.prefixes[i])
		}
	}
}

type removeAllValElementTest struct {
	testName    string
	inputArray  []int
	outputArray []int
	val         int
}

func TestRemoveAllValElements(t *testing.T) {
	testCases := []removeAllValElementTest{
		{"Case-1", []int{1, 2, 1, 4, 1}, []int{2, 4}, 1},
		{"Case-2", []int{}, []int{}, 1},
		{"Case-3", []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}, 1},
		{"Case-4", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, 5},
	}

	for _, testCase := range testCases {
		result := RemoveAllValElements(testCase.inputArray, testCase.val)
		slices.Sort(result)

		if !reflect.DeepEqual(result, testCase.outputArray) {
			t.Errorf("%s: %v != %v", testCase.testName, result, testCase.outputArray)
		}
	}
}

type FindSingleTest struct {
	testName    string
	inputArray  []int
	expectedRes int
}

func TestFindSingle(t *testing.T) {
	testCases := []FindSingleTest{
		{"Case-1", []int{1, 2, 1, 4, 2}, 4},
		{"Case-2", []int{}, -1},
		{"Case-3", []int{4, 1, 2, 1, 2}, 4},
		{"Case-4", []int{2, 2, 1}, 1},
		{"Case-5", []int{1}, 1},
	}

	for _, testCase := range testCases {
		result := FindSingle(testCase.inputArray)
		if result != testCase.expectedRes {
			t.Errorf("%s: %v != %v", testCase.testName, result, testCase.expectedRes)
		}
	}
}

type TestCaseIsPalindromeNumber struct {
	testName string
	x        int
	result   bool
}

func TestIsPalindromeNumber(t *testing.T) {
	testCases := []TestCaseIsPalindromeNumber{
		{"Case-1", 1, true},
		{"Case-2", 111, true},
		{"Case-3", 121, true},
		{"Case-4", 1221, true},
		{"Case-5", 9, true},
		{"Case-6", 10, false},
	}

	for _, testCase := range testCases {
		result := IsPalindromeNumber(testCase.x)
		if result != testCase.result {
			t.Errorf("%s: %v != %v", testCase.testName, result, testCase.result)
		}
	}
}

type testSearchInsertIndexReturn struct {
	testName    string
	inputArray  []int
	target      int
	expectedRes int
}

func TestSearchInsertIndex(t *testing.T) {
	testCases := []testSearchInsertIndexReturn{
		{"Case-1", []int{1, 3, 5, 6}, 5, 2},
		{"Case-2", []int{1, 3, 5, 6}, 2, 1},
		{"Case-3", []int{1, 3, 5, 6}, 7, 4},
		{"Case-4", []int{1}, 1, 0},
		{"Case-5", []int{1}, 0, 0},
		{"Case-6", []int{1, 2, 3}, 2, 1},
		{"Case-7", []int{1}, 2, 1},
		{"Case-8", []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20}, 9, 8},
	}

	for _, testCase := range testCases {
		res := SearchInsertIndexReturn(testCase.inputArray, testCase.target)
		if testCase.expectedRes != res {
			t.Errorf("%s: %v != %v", testCase.testName, res, testCase.expectedRes)
		}
	}
}
