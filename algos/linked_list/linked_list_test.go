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
			{2, 4, 3}, {9, 9, 9, 9, 9, 9, 9}, {0},
		},
		lists2: [][]int{
			{5, 6, 4}, {9, 9, 9, 9}, {0},
		},
		results: []string{
			"[7,0,8]", "[8,9,9,9,0,0,0,1]", "[0]",
		},
	}

	for i := 0; i < len(vars.lists1); i++ {
		node1 := &ListNode{vars.lists1[i][0], nil}
		if len(vars.lists1[i]) > 1 {
			node1.AddToTail(vars.lists1[i][1:]...)
		}

		node2 := &ListNode{vars.lists2[i][0], nil}
		if len(vars.lists2[i]) > 1 {
			node2.AddToTail(vars.lists2[i][1:]...)
		}

		result := AddTwoNumbers(node1, node2)
		resultStr := result.PrintList()

		if resultStr != vars.results[i] {
			fmt.Printf("%v != %v\n", resultStr, vars.results[i])
		}

	}
}
