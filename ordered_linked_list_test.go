package algos_in_golang

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
