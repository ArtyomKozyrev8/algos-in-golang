package algos_in_golang

import "testing"

func TestAddNode(t *testing.T) {
	list := LinkedList{}
	list.AddNode(1)

	if list.head == nil {
		t.Errorf("head of LinkedList is nil")
	}

	if list.head.value != 1 {
		t.Errorf("value of head node in LinkedList is not 1")
	}

	res := list.PrintList()
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
		list.AddNode(nodes[i])
		if list.PrintList() != expectedResults[i] {
			t.Errorf("%s != %s", list.PrintList(), expectedResults[i])
		}
	}
}

func TestReverseList(t *testing.T) {
	list := LinkedList{}
	list.ReverseList()
	if list.head != nil {
		t.Errorf("head of LinkedList is not nil")
	}

	res := list.PrintList()
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
		list.AddNode(nodes[i])
		list.ReverseList()
		if list.PrintList() != expectedResults[i] {
			t.Errorf("%s != %s", list.PrintList(), expectedResults[i])
		}
		if list.head.value != expectedHeadValue[i] {
			t.Errorf("%d != %d", list.head.value, expectedHeadValue[i])
		}
	}
}

func TestGetNodeByIndex(t *testing.T) {
	list := LinkedList{}
	if list.GetNodeByIndex(0) != nil {
		t.Errorf("GetNodeByIndex(0) is not nil")
	}

	list.AddNode(1)
	if list.GetNodeByIndex(0).value != 1 {
		t.Errorf("GetNodeByIndex(0) is not 1")
	}

	list.AddNode(2)
	if list.GetNodeByIndex(0).value != 1 {
		t.Errorf("GetNodeByIndex(0) is not 1")
	}
	if list.GetNodeByIndex(1).value != 2 {
		t.Errorf("GetNodeByIndex(1) is not 2")
	}

	nodes := []int{
		3, 4, 5,
	}
	for _, node := range nodes {
		list.AddNode(node)
	}

	expectedResults := []int{1, 2, 3, 4, 5}

	for index, value := range expectedResults {
		if list.GetNodeByIndex(index).value != value {
			t.Errorf("GetNodeByIndex(%d) is not %d", index, list.GetNodeByIndex(index).value)
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

	list.AddNode(1)
	if list.GetNodeByValue(1).value != 1 {
		t.Errorf("GetNodeByValue(1) is not 1")
	}

	list.AddNode(2)
	if list.GetNodeByValue(2).value != 2 {
		t.Errorf("GetNodeByValue(2) is not 2")
	}

	nodes := []int{3, 4, 5}
	for _, node := range nodes {
		list.AddNode(node)
	}

	nodes = []int{1, 2, 3, 4, 5}

	for _, value := range nodes {
		if list.GetNodeByValue(value).value != value {
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
	if list.head.value != 1 {
		t.Errorf("value of head node in LinkedList is not 1")
	}
	list.AddNodeToHead(2)
	if list.head.value != 2 {
		t.Errorf("value of head node in LinkedList is not 2")
	}
	resStr := list.PrintList()
	if resStr != "LinkedList: 2,1" {
		t.Errorf("%s != LinkedList: 2,1", resStr)
	}
	list.AddNodeToHead(3)
	if list.head.value != 3 {
		t.Errorf("value of head node in LinkedList is not 3")
	}
	resStr = list.PrintList()
	if resStr != "LinkedList: 3,2,1" {
		t.Errorf("%s != LinkedList: 3,2,1", resStr)
	}
	list.AddNode(4)
	resStr = list.PrintList()
	if resStr != "LinkedList: 3,2,1,4" {
		t.Errorf("%s != LinkedList: 3,2,1,4", resStr)
	}
}
