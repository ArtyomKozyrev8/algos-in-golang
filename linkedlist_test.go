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
