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
