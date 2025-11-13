package linkedlist

import (
	"fmt"
	"testing"
)

func TestAddNodeToTail(t *testing.T) {
	list := LinkedList{}
	list.AddNodeToTail(1)

	if list.head == nil {
		t.Errorf("head of LinkedList is nil")
	}

	if list.head.Value != 1 {
		t.Errorf("value of head node in LinkedList is not 1")
	}

	res := fmt.Sprintf("%v", &list)
	if res != "LinkedList: 1" {
		t.Errorf("%s != LinkedList: 1", res)
	}

	nodes := []int{
		2, 3, 4, 5,
	}
	expectedResults := []string{
		"LinkedList: 1,2", "LinkedList: 1,2,3", "LinkedList: 1,2,3,4", "LinkedList: 1,2,3,4,5",
	}

	for i := 0; i < len(nodes); i++ {
		list.AddNodeToTail(nodes[i])
		if fmt.Sprintf("%v", &list) != expectedResults[i] {
			t.Errorf("%s != %s", fmt.Sprintf("%v", &list), expectedResults[i])
		}
	}
}

func TestReverseList(t *testing.T) {
	list := LinkedList{}
	list.ReverseList()
	if list.head != nil {
		t.Errorf("head of LinkedList is not nil")
	}

	res := fmt.Sprintf("%v", &list)
	if res != "LinkedList: " {
		t.Errorf("%s != LinkedList: ", res)
	}

	nodes := []int{
		1, 2, 3, 4, 5,
	}
	expectedResults := []string{
		"LinkedList: 1", "LinkedList: 2,1", "LinkedList: 3,1,2", "LinkedList: 4,2,1,3", "LinkedList: 5,3,1,2,4",
	}
	expectedHeadValue := []int{
		1, 2, 3, 4, 5,
	}

	for i := 0; i < len(nodes); i++ {
		list.AddNodeToTail(nodes[i])
		list.ReverseList()
		if fmt.Sprintf("%v", &list) != expectedResults[i] {
			t.Errorf("%s != %s", fmt.Sprintf("%v", &list), expectedResults[i])
		}
		if list.head.Value != expectedHeadValue[i] {
			t.Errorf("%d != %d", list.head.Value, expectedHeadValue[i])
		}
	}
}

func TestGetNodeByIndex(t *testing.T) {
	list := LinkedList{}
	if list.GetNodeByIndex(0) != nil {
		t.Errorf("GetNodeByIndex(0) is not nil")
	}

	list.AddNodeToTail(1)
	if list.GetNodeByIndex(0).Value != 1 {
		t.Errorf("GetNodeByIndex(0) is not 1")
	}

	list.AddNodeToTail(2)
	if list.GetNodeByIndex(0).Value != 1 {
		t.Errorf("GetNodeByIndex(0) is not 1")
	}
	if list.GetNodeByIndex(1).Value != 2 {
		t.Errorf("GetNodeByIndex(1) is not 2")
	}

	nodes := []int{
		3, 4, 5,
	}
	for _, node := range nodes {
		list.AddNodeToTail(node)
	}

	expectedResults := []int{1, 2, 3, 4, 5}

	for index, value := range expectedResults {
		if list.GetNodeByIndex(index).Value != value {
			t.Errorf("GetNodeByIndex(%d) is not %d", index, list.GetNodeByIndex(index).Value)
		}
	}

	if list.GetNodeByIndex(10) != nil {
		t.Errorf("GetNodeByIndex(10) is not nil")
	}
}

func TestGetNodeByValue(t *testing.T) {
	list := LinkedList{}
	if list.GetNodeByValue(10) != nil {
		t.Errorf("GetNodeByValue(10) is not nil")
	}

	list.AddNodeToTail(1)
	if list.GetNodeByValue(1).Value != 1 {
		t.Errorf("GetNodeByValue(1) is not 1")
	}

	list.AddNodeToTail(2)
	if list.GetNodeByValue(2).Value != 2 {
		t.Errorf("GetNodeByValue(2) is not 2")
	}

	nodes := []int{3, 4, 5}
	for _, node := range nodes {
		list.AddNodeToTail(node)
	}

	nodes = []int{1, 2, 3, 4, 5}

	for _, value := range nodes {
		if list.GetNodeByValue(value).Value != value {
			t.Errorf("GetNodeByValue(%d) is not %d", value, value)
		}
	}

	if list.GetNodeByValue(10) != nil {
		t.Errorf("GetNodeByIndex(10) is not nil")
	}
}

func TestAddNodeToHead(t *testing.T) {
	list := LinkedList{}
	list.AddNodeToHead(1)
	if list.head == nil {
		t.Errorf("head of LinkedList is nil")
	}
	if list.head.Value != 1 {
		t.Errorf("value of head node in LinkedList is not 1")
	}
	list.AddNodeToHead(2)
	if list.head.Value != 2 {
		t.Errorf("value of head node in LinkedList is not 2")
	}
	resStr := fmt.Sprintf("%v", &list)
	if resStr != "LinkedList: 2,1" {
		t.Errorf("%s != LinkedList: 2,1", resStr)
	}
	list.AddNodeToHead(3)
	if list.head.Value != 3 {
		t.Errorf("value of head node in LinkedList is not 3")
	}
	resStr = fmt.Sprintf("%v", &list)
	if resStr != "LinkedList: 3,2,1" {
		t.Errorf("%s != LinkedList: 3,2,1", resStr)
	}
	list.AddNodeToTail(4)
	resStr = fmt.Sprintf("%v", &list)
	if resStr != "LinkedList: 3,2,1,4" {
		t.Errorf("%s != LinkedList: 3,2,1,4", resStr)
	}
}

func TestRemoveNodeFromHead(t *testing.T) {
	list := LinkedList{}
	if list.RemoveNodeFromHead() != nil {
		t.Errorf("RemoveNodeFromHead() is not nil")
	}
	if list.head != nil {
		t.Errorf("head of LinkedList is not nil")
	}

	list.AddNodeToHead(1)
	if list.RemoveNodeFromHead().Value != 1 {
		t.Errorf("remove node value is not 1")
	}
	if list.head != nil {
		t.Errorf("head of LinkedList is not nil")
	}

	list.AddNodeToTail(1)
	list.AddNodeToTail(2)
	if list.RemoveNodeFromHead().Value != 1 {
		t.Errorf("remove node value is not 1")
	}

	list.AddNodeToTail(1)
	list.AddNodeToTail(3)
	if list.RemoveNodeFromHead().Value != 2 {
		t.Errorf("remove node value is not 2")
	}
	resStr := fmt.Sprintf("%v", &list)
	if resStr != "LinkedList: 1,3" {
		t.Errorf("%s != LinkedList: 1,3", resStr)
	}

	list.AddNodeToTail(4)
	list.AddNodeToHead(5)
	removedNode := list.RemoveNodeFromHead()
	if removedNode == nil {
		t.Errorf("removedNode is nil")
	} else {
		if removedNode.Value != 5 {
			t.Errorf("removedNode value is not 5")
		}
		if removedNode.Next.Value != 1 {
			t.Errorf("removedNode next node value is not 1")
		}
		resStr = fmt.Sprintf("%v", &list)
		if resStr != "LinkedList: 1,3,4" {
			t.Errorf("%s != LinkedList: 1,3,4", resStr)
		}
	}
}

func TestRemoveNodeFromTail(t *testing.T) {
	list := LinkedList{}
	list.RemoveNodeFromTail()
	if list.head != nil {
		t.Errorf("head of LinkedList is not nil")
	}

	list.AddNodeToTail(1)
	removedNode := list.RemoveNodeFromTail()
	if removedNode == nil {
		t.Errorf("removedNode is nil")
	} else {
		if removedNode.Value != 1 {
			t.Errorf("removedNode value is not 1")
		}
		if list.head != nil {
			t.Errorf("head of LinkedList is not nil")
		}
	}

	list.AddNodeToTail(1)
	list.AddNodeToTail(2)
	removedNode = list.RemoveNodeFromTail()
	if removedNode == nil {
		t.Errorf("removedNode is nil")
	} else {
		if removedNode.Value != 2 {
			t.Errorf("removedNode value is not 2")
		}
	}

	list.AddNodeToTail(2)
	list.AddNodeToTail(3)
	list.AddNodeToHead(4)
	removedNode = list.RemoveNodeFromTail()
	if removedNode == nil {
		t.Errorf("removedNode is nil")
	} else {
		if removedNode.Value != 3 {
			t.Errorf("removedNode value is not3")
		}
	}
	listStr := fmt.Sprintf("%v", &list)
	if listStr != "LinkedList: 4,1,2" {
		t.Errorf("%s != LinkedList: 4,1,2", listStr)
	}
}

func TestLength(t *testing.T) {
	list := LinkedList{}
	if list.Length() != 0 {
		t.Errorf("length of LinkedList is not 0")
	}

	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		list.AddNodeToTail(value)
		if list.Length() != value {
			t.Errorf("length of LinkedList is not %d", value)
		}
	}

	list.AddNodeToHead(6)
	if list.Length() != 6 {
		t.Errorf("length of LinkedList is not 6")
	}
	listStr := fmt.Sprintf("%v", &list)
	if listStr != "LinkedList: 6,1,2,3,4,5" {
		t.Errorf("%s != LinkedList: 6,1,2,3,4,5", listStr)
	}
}

func TestRemoveNodeByIndex(t *testing.T) {
	list := LinkedList{}
	err, removedNode := list.RemoveNodeByIndex(0)
	if err == nil || removedNode != nil {
		t.Errorf("RemoveNodeByIndex(0) is not nil or did not raise an error")
	}

	list.AddNodeToTail(1)
	err, removedNode = list.RemoveNodeByIndex(-1)
	if err == nil || removedNode != nil {
		t.Errorf("RemoveNodeByIndex(-1) is not nil or did not raise an error")
	}

	err, removedNode = list.RemoveNodeByIndex(1)
	if err == nil || removedNode != nil {
		t.Errorf("RemoveNodeByIndex(1) is not nil or did not raise an error")
	}

	err, removedNode = list.RemoveNodeByIndex(0)
	if err != nil || removedNode.Value != 1 || fmt.Sprintf("%v", &list) != "LinkedList: " {
		t.Errorf("RemoveNodeByIndex(0) is not nil or did not raise an error or not equal 1")
	}

	nodes := [][]int{
		{1, 2},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	indexToRemoves := []int{
		0,
		1,
		0,
		1,
		2,
		0,
		1,
		2,
		3,
		0,
		1,
		2,
		3,
		4,
	}
	expectedRemoveNodeValues := []int{
		1,
		2,
		1,
		2,
		3,
		1,
		2,
		3,
		4,
		1,
		2,
		3,
		4,
		5,
	}
	expectedStrRes := []string{
		"LinkedList: 2",
		"LinkedList: 1",
		"LinkedList: 2,3",
		"LinkedList: 1,3",
		"LinkedList: 1,2",
		"LinkedList: 2,3,4",
		"LinkedList: 1,3,4",
		"LinkedList: 1,2,4",
		"LinkedList: 1,2,3",
		"LinkedList: 2,3,4,5",
		"LinkedList: 1,3,4,5",
		"LinkedList: 1,2,4,5",
		"LinkedList: 1,2,3,5",
		"LinkedList: 1,2,3,4",
	}

	for i := 0; i < len(nodes); i++ {
		list = LinkedList{}
		for j := 0; j < len(nodes[i]); j++ {
			list.AddNodeToTail(nodes[i][j])
		}
		_, removedNodeRes := list.RemoveNodeByIndex(indexToRemoves[i])
		if removedNodeRes == nil {
			t.Errorf("RemoveNodeByIndex(%d) is nil", indexToRemoves[i])
		} else {
			if removedNodeRes.Value != expectedRemoveNodeValues[i] {
				t.Errorf(
					"RemoveNodeByIndex(%d) value %d is not equal to %d",
					indexToRemoves[i],
					removedNodeRes.Value,
					expectedRemoveNodeValues[i],
				)
			}
			if fmt.Sprintf("%v", &list) != expectedStrRes[i] {
				t.Errorf(
					"RemoveNodeByIndex(%d) is not equal to %s",
					indexToRemoves[i],
					expectedStrRes[i],
				)
			}
		}
	}
}

func TestSortWayOneAsc(t *testing.T) {
	input := struct {
		nodes          [][]int
		expectedValues []string
	}{
		[][]int{
			{},
			{1},
			{1, 0},
			{0, 1},
			{0, 1, 2},
			{2, 0, 1},
			{2, 1, 0},
			{4, 4, 0, 2, 1, 3, 9},
			{0, 0, 0, 2, 1, 3, 9, 4, 3, 5, 6, 1, 2, 5},
			{3, 4, 1, 3, 2, 5, 1, 6, 7, 1, 7, 1, 0, 3, 4, 3, 5, 6, 7, 8, 9, 4, 4, 2, 1, 0},
		},
		[]string{
			"LinkedList: ",
			"LinkedList: 1",
			"LinkedList: 0,1",
			"LinkedList: 0,1",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2,3,4,4,9",
			"LinkedList: 0,0,0,1,1,2,2,3,3,4,5,5,6,9",
			"LinkedList: 0,0,1,1,1,1,1,2,2,3,3,3,3,4,4,4,4,5,5,6,6,7,7,7,8,9",
		},
	}

	for i := 0; i < len(input.nodes); i++ {
		list := LinkedList{}
		for j := 0; j < len(input.nodes[i]); j++ {
			list.AddNodeToTail(input.nodes[i][j])
		}
		list.SortWayOne(true)
		resStr := fmt.Sprintf("%v", &list)
		if resStr != input.expectedValues[i] {
			t.Errorf("%s != %s", resStr, input.expectedValues[i])
		}
	}
}

func TestSortWayTwoAsc(t *testing.T) {
	input := struct {
		nodes          [][]int
		expectedValues []string
	}{
		[][]int{
			{},
			{1},
			{1, 0},
			{0, 1},
			{0, 1, 2},
			{2, 0, 1},
			{2, 1, 0},
			{4, 4, 0, 2, 1, 3, 9},
			{0, 0, 0, 2, 1, 3, 9, 4, 3, 5, 6, 1, 2, 5},
			{3, 4, 1, 3, 2, 5, 1, 6, 7, 1, 7, 1, 0, 3, 4, 3, 5, 6, 7, 8, 9, 4, 4, 2, 1, 0},
		},
		[]string{
			"LinkedList: ",
			"LinkedList: 1",
			"LinkedList: 0,1",
			"LinkedList: 0,1",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2",
			"LinkedList: 0,1,2,3,4,4,9",
			"LinkedList: 0,0,0,1,1,2,2,3,3,4,5,5,6,9",
			"LinkedList: 0,0,1,1,1,1,1,2,2,3,3,3,3,4,4,4,4,5,5,6,6,7,7,7,8,9",
		},
	}

	for i := 0; i < len(input.nodes); i++ {
		list := LinkedList{}
		for j := 0; j < len(input.nodes[i]); j++ {
			list.AddNodeToTail(input.nodes[i][j])
		}
		list.SortWayTwo(true)
		resStr := fmt.Sprintf("%v", &list)
		if resStr != input.expectedValues[i] {
			t.Errorf("%s != %s", resStr, input.expectedValues[i])
		}
	}
}

func TestSortWayTwoDesc(t *testing.T) {
	input := struct {
		nodes          [][]int
		expectedValues []string
	}{
		[][]int{
			{},
			{1},
			{1, 0},
			{0, 1},
			{0, 1, 2},
			{2, 0, 1},
			{2, 1, 0},
			{4, 4, 0, 2, 1, 3, 9},
			{0, 0, 0, 2, 1, 3, 9, 4, 3, 5, 6, 1, 2, 5},
			{3, 4, 1, 3, 2, 5, 1, 6, 7, 1, 7, 1, 0, 3, 4, 3, 5, 6, 7, 8, 9, 4, 4, 2, 1, 0},
		},
		[]string{
			"LinkedList: ",
			"LinkedList: 1",
			"LinkedList: 1,0",
			"LinkedList: 1,0",
			"LinkedList: 2,1,0",
			"LinkedList: 2,1,0",
			"LinkedList: 2,1,0",
			"LinkedList: 9,4,4,3,2,1,0",
			"LinkedList: 9,6,5,5,4,3,3,2,2,1,1,0,0,0",
			"LinkedList: 9,8,7,7,7,6,6,5,5,4,4,4,4,3,3,3,3,2,2,1,1,1,1,1,0,0",
		},
	}

	for i := 0; i < len(input.nodes); i++ {
		list := LinkedList{}
		for j := 0; j < len(input.nodes[i]); j++ {
			list.AddNodeToTail(input.nodes[i][j])
		}
		list.SortWayTwo(false)
		resStr := fmt.Sprintf("%v", &list)
		if resStr != input.expectedValues[i] {
			t.Errorf("%s != %s", resStr, input.expectedValues[i])
		}
	}
}
