package linked_list

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbersTest(t *testing.T) {
	vars := struct {
		lists1  [][]int
		lists2  [][]int
		results []string
	}{
		lists1: [][]int{
			{2, 4, 3}, {9, 9, 9, 9, 9, 9, 9}, {0}, {9, 9, 9, 9},
		},
		lists2: [][]int{
			{5, 6, 4}, {9, 9, 9, 9}, {0}, {9, 9, 9, 9, 9, 9, 9},
		},
		results: []string{
			"[7,0,8]", "[8,9,9,9,0,0,0,1]", "[0]", "[8,9,9,9,0,0,0,1]",
		},
	}

	for i := 0; i < len(vars.lists1); i++ {
		node1 := BuildList(vars.lists1[i]...)
		node2 := BuildList(vars.lists2[i]...)
		result := AddTwoNumbers(node1, node2)
		resultStr := result.PrintList()

		if resultStr != vars.results[i] {
			t.Errorf("%v != %v\n", resultStr, vars.results[i])
		}

	}
}

func TestRotateRight(t *testing.T) {
	vars := struct {
		lists   [][]int
		k       []int
		results []string
	}{
		lists: [][]int{
			{0, 1, 2}, {1, 2, 3, 4, 5}, {1, 2, 3}, {0},
		},
		k: []int{4, 2, 2_000_000_000, 2_000_000_000},
		results: []string{
			"[2,0,1]", "[4,5,1,2,3]", "[2,3,1]", "[0]",
		},
	}

	for i := 0; i < len(vars.lists); i++ {
		node := BuildList(vars.lists[i]...)

		result := RotateRight(node, vars.k[i])
		resultStr := result.PrintList()

		if resultStr != vars.results[i] {
			t.Errorf("%v != %v\n", resultStr, vars.results[i])
		}
	}
}

func TestRotateRightEmpty(t *testing.T) {
	var node *ListNode
	result := RotateRight(node, 1000)
	if result != nil {
		fmt.Printf("%v != %v\n", result, nil)
	}

	if node.PrintList() != "[]" {
		t.Errorf("%v != %v\n", node.PrintList(), "[]")
	}
}

func TestAddToTail(t *testing.T) {
	var node *ListNode = &ListNode{0, nil}
	node.AddToTail(1)
	node.AddToTail(1, 2, 3)

	if node.PrintList() != "[0,1,1,2,3]" {
		t.Errorf("%v != %v\n", node.PrintList(), "[0,1,1,2,3]")
	}
}

func TestMergeTwoSortedListAsSortedList(t *testing.T) {
	vars := struct {
		caseNames []string
		lists1    [][]int
		lists2    [][]int
		results   []string
	}{
		caseNames: []string{
			"Case1", "Case2", "Case3", "Case4", "Case5", "Case6", "Case7", "Case8", "Case9"},
		lists1: [][]int{
			{1, 2, 4}, {}, {}, {0}, {1, 2, 4}, {2, 4, 4}, {4, 5, 6}, {1, 2, 3}, {},
		},
		lists2: [][]int{
			{1, 3, 4}, {}, {0}, {}, {2, 3, 4}, {1}, {1, 2, 3}, {4, 5, 6}, {1, 2, 3},
		},
		results: []string{
			"[1,1,2,3,4,4]", "[]", "[0]", "[0]", "[1,2,2,3,4,4]", "[1,2,4,4]", "[1,2,3,4,5,6]", "[1,2,3,4,5,6]", "[1,2,3]",
		},
	}

	for i := 0; i < len(vars.lists1); i++ {
		var node1 *ListNode
		var node2 *ListNode
		if len(vars.lists1[i]) > 0 {
			node1 = BuildList(vars.lists1[i]...)
		}

		if len(vars.lists2[i]) > 0 {
			node2 = BuildList(vars.lists2[i]...)
		}

		res := MergeTwoSortedListAsSortedList(node1, node2)
		resStr := res.PrintList()
		if resStr != vars.results[i] {
			t.Errorf("%s: %v != %v\n", vars.caseNames[i], resStr, vars.results[i])
		}
	}
}

type TestCasesForFindMiddleNode struct {
	nodeVals       []int
	expectedResult int
}

func TestFindMiddleNode(t *testing.T) {
	tests := []TestCasesForFindMiddleNode{
		{[]int{1, 2, 3, 4, 5, 6}, 4}, {[]int{1, 2, 3, 4, 5}, 3}, {[]int{}, -1},
	}

	for _, test := range tests {
		list := BuildList(test.nodeVals...)
		res := FindMiddleNodeValue(list)
		if res != test.expectedResult {
			t.Errorf("%v != %v\n", res, test.expectedResult)
		}
	}
}

type RemoveNthFromEndTest struct {
	caseName       string
	nodeVals       []int
	val            int
	expectedResult string
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []RemoveNthFromEndTest{
		{"Case-1", []int{1, 2}, 2, "[2]"},
		{"Case-2", []int{1, 2, 3, 4, 5}, 2, "[1,2,3,5]"},
		{"Case-3", []int{1}, 1, "[]"},
		{"Case-4", []int{1, 2}, 1, "[1]"},
	}
	for _, test := range tests {
		list := BuildList(test.nodeVals...)
		result := RemoveNthFromEnd(list, test.val)
		resultStr := result.PrintList()
		if resultStr != test.expectedResult {
			t.Errorf("%s: %v != %v\n", test.caseName, resultStr, test.expectedResult)
		}
	}
}

type TestCaseSwapTwoNearest struct {
	caseName       string
	nodes          []int
	expectedResult string
}

func TestSwapTwoNearest(t *testing.T) {
	tests := []TestCaseSwapTwoNearest{
		{"Case-1", []int{1}, "[1]"},
		{"Case-2", []int{}, "[]"},
		{"Case-3", []int{1, 2}, "[2,1]"},
		{"Case-4", []int{1, 2, 3}, "[2,1,3]"},
		{"Case-5", []int{1, 2, 3, 4}, "[2,1,4,3]"},
		{"Case-6", []int{1, 2, 3, 4, 5}, "[2,1,4,3,5]"},
	}

	for _, test := range tests {
		list := BuildList(test.nodes...)
		res := SwapTwoNearest(list)
		resStr := res.PrintList()
		if resStr != test.expectedResult {
			t.Errorf("%s: %v != %v", test.caseName, resStr, test.expectedResult)
		}
	}
}

type TestCaseReverseList struct {
	caseName       string
	nodes          []int
	expectedResult string
}

func TestReverseList(t *testing.T) {
	tests := []TestCaseReverseList{
		{"Case-0", []int{}, "[]"},
		{"Case-1", []int{1}, "[1]"},
		{"Case-2", []int{1, 2}, "[2,1]"},
		{"Case-3", []int{1, 2, 3}, "[3,2,1]"},
		{"Case-4", []int{1, 2, 3, 4}, "[4,3,2,1]"},
		{"Case-5", []int{1, 2, 3, 4, 5}, "[5,4,3,2,1]"},
	}

	for _, test := range tests {
		list := BuildList(test.nodes...)
		res := ReverseList(list)
		resStr := res.PrintList()
		if resStr != test.expectedResult {
			t.Errorf("%s: %v != %v", test.caseName, resStr, test.expectedResult)
		}
	}
}

type testDeleteDuplicates struct {
	caseName       string
	nodes          []int
	expectedResult string
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []testDeleteDuplicates{
		{"Case-1", []int{1}, "[1]"},
		{"Case-2", []int{1, 1, 1, 2}, "[1,2]"},
		{"Case-3", []int{1, 1, 2, 2, 2, 2, 2, 3, 3, 3}, "[1,2,3]"},
	}

	for _, test := range tests {
		list := BuildList(test.nodes...)
		res := DeleteDuplicates(list)
		resStr := res.PrintList()
		if resStr != test.expectedResult {
			t.Errorf("%s: %v != %v", test.caseName, resStr, test.expectedResult)
		}
	}
}
