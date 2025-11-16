package ordered_linked_list

import (
	"fmt"
	"testing"
)

func TestOrderedAscLinkedList_String(t *testing.T) {
	list := OrderedAscLinkedList{}
	resStr := fmt.Sprintf("%v", &list)
	if resStr != "OAscLL: " {
		t.Error("Expected 'OAscLL ' got", resStr)
	}

	list.Append(0)
	resStr = fmt.Sprintf("%v", &list)
	if resStr != "OAscLL: 0" {
		t.Error("Expected 'OAscLL 0' got", resStr)
	}

	vars := struct {
		values  []int
		results []string
	}{
		values:  []int{1, 2, 3, 4, 5},
		results: []string{"OAscLL: 0,1", "OAscLL: 0,1,2", "OAscLL: 0,1,2,3", "OAscLL: 0,1,2,3,4", "OAscLL: 0,1,2,3,4,5"},
	}

	for i := 0; i < len(vars.values); i++ {
		list.Append(vars.values[i])
		resStr = fmt.Sprintf("%v", &list)
		if resStr != vars.results[i] {
			t.Error("Expected ", vars.results[i], "got", resStr)
		}
	}
}

func TestOrderedAscLinkedList_Append(t *testing.T) {
	vars := struct {
		values  [][]int
		results []string
	}{
		values: [][]int{
			{1, 10},
			{10, 1},
			{10, 2, 3},
			{2, 10, 3},
			{2, 3, 10},
			{10, 3, 2},
			{3, 2, 10},
			{3, 10, 2},
			{1, 2, 3, 4},
			{2, 1, 3, 4},
			{2, 3, 1, 4},
			{2, 3, 4, 1},
			{2, 1, 3, 4},
			{2, 3, 1, 4},
			{2, 3, 4, 1},
			{3, 1, 2, 4},
			{1, 3, 2, 4},
			{1, 2, 4, 3},
			{4, 1, 2, 3},
			{1, 4, 2, 3},
			{1, 2, 4, 3},
		},
		results: []string{
			"OAscLL: 1,10",
			"OAscLL: 1,10",
			"OAscLL: 2,3,10",
			"OAscLL: 2,3,10",
			"OAscLL: 2,3,10",
			"OAscLL: 2,3,10",
			"OAscLL: 2,3,10",
			"OAscLL: 2,3,10",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,4",
		},
	}

	for i := 0; i < len(vars.values); i++ {
		list := OrderedAscLinkedList{}
		for j := 0; j < len(vars.values[i]); j++ {
			list.Append(vars.values[i][j])
		}
		resStr := fmt.Sprintf("%v", &list)
		if resStr != vars.results[i] {
			t.Error("Expected ", vars.results[i], "got", resStr)
		}
	}
}

func TestOrderedAscLinkedList_RemoveByValue(t *testing.T) {
	vars := struct {
		values      [][]int
		removeValue []int
		fineResult  []int
		errResults  []error
		resStr      []string
	}{
		values: [][]int{
			{}, {3}, {3}, {2, 1}, {1, 2}, {1, 2}, {3, 4, 5}, {4, 3, 5}, {5, 3, 4}, {3, 4, 5}, {1, 2, 5, 6, 7, 8},
		},
		removeValue: []int{
			1, 3, 9, 1, 2, 9, 3, 4, 5, 9, 3,
		},
		fineResult: []int{
			0, 3, 0, 1, 2, 0, 3, 4, 5, 0, 0,
		},
		errResults: []error{
			&OLLError{"value 1 not found"},
			nil,
			&OLLError{"value 9 not found"},
			nil,
			nil,
			&OLLError{"value 9 not found"},
			nil,
			nil,
			nil,
			&OLLError{"value 9 not found"},
			&OLLError{"value 3 not found"},
		},
		resStr: []string{
			"OAscLL: ",
			"OAscLL: ",
			"OAscLL: 3",
			"OAscLL: 2",
			"OAscLL: 1",
			"OAscLL: 1,2",
			"OAscLL: 4,5",
			"OAscLL: 3,5",
			"OAscLL: 3,4",
			"OAscLL: 3,4,5",
			"OAscLL: 1,2,5,6,7,8",
		},
	}
	for i := 0; i < len(vars.values); i++ {
		list := OrderedAscLinkedList{}
		for j := 0; j < len(vars.values[i]); j++ {
			list.Append(vars.values[i][j])
		}
		res, exc := list.RemoveByValue(vars.removeValue[i])

		if res != vars.fineResult[i] {
			t.Error("Expected ", vars.fineResult[i], "got", res)
		}

		if fmt.Sprintf("%v", exc) != fmt.Sprintf("%v", vars.errResults[i]) {
			t.Error(fmt.Sprintf("Expected '%v' != '%v'", vars.errResults[i], exc))
		}

		if fmt.Sprintf("%v", &list) != vars.resStr[i] {
			t.Error(fmt.Sprintf("Expected '%v' != '%v'", vars.resStr[i], &list))
		}
	}
}

func TestOrderedAscLinkedList_RemoveByIndex(t *testing.T) {
	vars := struct {
		values      [][]int
		indexRemove []int
		valReturn   []int
		errorReturn []error
		resStr      []string
	}{
		values: [][]int{
			{},
			{3},
			{3},
			{2, 1},
			{1, 2},
			{1, 2},
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
		indexRemove: []int{1, 0, 1, 0, 1, 5, 0, 1, 2, 5, 0, 4, 3, 9},
		valReturn:   []int{0, 3, 0, 1, 2, 0, 1, 2, 3, 0, 1, 5, 4, 0},
		errorReturn: []error{
			&OLLError{"index 1 out of range"},
			nil,
			&OLLError{"index 1 out of range"},
			nil,
			nil,
			&OLLError{"index 5 out of range"},
			nil,
			nil,
			nil,
			&OLLError{"index 5 out of range"},
			nil,
			nil,
			nil,
			&OLLError{"index 9 out of range"},
		},
		resStr: []string{
			"OAscLL: ",
			"OAscLL: ",
			"OAscLL: 3",
			"OAscLL: 2",
			"OAscLL: 1",
			"OAscLL: 1,2",
			"OAscLL: 2,3",
			"OAscLL: 1,3",
			"OAscLL: 1,2",
			"OAscLL: 1,2,3",
			"OAscLL: 2,3,4,5",
			"OAscLL: 1,2,3,4",
			"OAscLL: 1,2,3,5",
			"OAscLL: 1,2,3,4,5",
		},
	}

	for i := 0; i < len(vars.values); i++ {
		list := OrderedAscLinkedList{}
		for j := 0; j < len(vars.values[i]); j++ {
			list.Append(vars.values[i][j])
		}
		res, exc := list.RemoveByIndex(vars.indexRemove[i])
		if res != vars.valReturn[i] {
			t.Error("Expected ", vars.valReturn[i], "got", res)
		}
		if fmt.Sprintf("%v", exc) != fmt.Sprintf("%v", vars.errorReturn[i]) {
			t.Error(fmt.Sprintf("Expected '%v' != '%v'", vars.errorReturn[i], exc))
		}

		if fmt.Sprintf("%v", &list) != vars.resStr[i] {
			t.Error(fmt.Sprintf("Expected '%v' != '%v'", vars.resStr[i], &list))
		}
	}
}
