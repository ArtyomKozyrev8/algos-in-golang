package linkedlist

import (
	"errors"
	"strconv"

	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/nodes"
)

type LinkedList struct {
	head *nodes.Node
}

// AddNodeToTail adds new node to LinkedList
func (list *LinkedList) AddNodeToTail(value int) {
	if list.head == nil {
		list.head = &nodes.Node{Value: value, Next: nil}
	} else {
		cur := list.head
		for {
			if cur.Next == nil {
				cur.Next = &nodes.Node{Value: value, Next: nil}
				break
			}
			cur = cur.Next
		}
	}
}

// String converts nodes in LinkedList to string
func (list *LinkedList) String() string {
	res := "LinkedList: "
	cur := list.head

	for cur != nil {
		res += strconv.Itoa(cur.Value)
		if cur.Next != nil {
			res += ","
		}
		cur = cur.Next
	}

	return res
}

// ReverseList reverse order of nodes in LinkedList
func (list *LinkedList) ReverseList() {
	var prev *nodes.Node
	cur := list.head

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	list.head = prev
}

// GetNodeByIndex returns Node from LinkedList by index
func (list *LinkedList) GetNodeByIndex(index int) *nodes.Node {
	cur := list.head
	curIndex := 0
	for cur != nil {
		if curIndex == index {
			return cur
		}
		cur = cur.Next
		curIndex += 1
	}
	return nil
}

// GetNodeByValue returns Node from LinkedList by value
func (list *LinkedList) GetNodeByValue(value int) *nodes.Node {
	cur := list.head
	for cur != nil {
		if cur.Value == value {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// AddNodeToHead adds new node to the head of LinkedList
func (list *LinkedList) AddNodeToHead(value int) {
	if list.head == nil {
		list.head = &nodes.Node{Value: value, Next: nil}
	} else {
		cur := list.head
		list.head = &nodes.Node{Value: value, Next: cur}
	}
}

// RemoveNodeFromHead removes and returns Node from head of LinkedList
func (list *LinkedList) RemoveNodeFromHead() *nodes.Node {
	if list.head == nil {
		return nil
	}

	cur := list.head
	curNext := cur.Next
	list.head = curNext

	return cur
}

// RemoveNodeFromTail removes and returns Node from tail of LinkedList
func (list *LinkedList) RemoveNodeFromTail() *nodes.Node {
	if list.head == nil {
		return nil
	}

	var prev *nodes.Node
	cur := list.head
	for {
		if cur.Next != nil {
			prev = cur
			cur = cur.Next
		} else {
			if prev == nil {
				list.head = nil
			} else {
				prev.Next = nil
			}
			return cur
		}
	}
}

// Length returns number of Nodes in LinkedList
func (list *LinkedList) Length() int {
	num := 0
	cur := list.head

	for cur != nil {
		cur = cur.Next
		num += 1
	}

	return num
}

// RemoveNodeByIndex removes Node from LinkedList based on index
func (list *LinkedList) RemoveNodeByIndex(index int) (error, *nodes.Node) {
	if index < 0 {
		return errors.New("index can't be negative number"), nil
	}

	if list.head == nil {
		return errors.New("can't remove from empty list"), nil
	}

	curIndex := 0
	var prev *nodes.Node
	cur := list.head

	for cur != nil {
		if curIndex == index {
			if prev == nil {
				list.head = cur.Next
			} else {
				prev.Next = cur.Next
			}
			return nil, cur
		} else {
			prev = cur
			cur = cur.Next
			curIndex += 1
		}
	}

	return errors.New("index out of range"), nil
}

// SortWayOne - my first idea how to sort LinkedList
func (list *LinkedList) SortWayOne(orderAsc bool) {
	sortedList := LinkedList{}

	for list.head != nil {
		cur := list.head
		curIndex := 0
		curExtremumIndex := 0
		curExtremumValue := list.head.Value
		for cur != nil {
			if cur.Next == nil {
				break
			} else {
				curIndex += 1
				if orderAsc == true {
					if curExtremumValue <= cur.Next.Value {
						curExtremumValue = cur.Next.Value
						curExtremumIndex = curIndex
					}
				} else {
					if curExtremumValue >= cur.Next.Value {
						curExtremumValue = cur.Next.Value
						curExtremumIndex = curIndex
					}
				}
				cur = cur.Next
			}
		}
		_, nodeExtremum := list.RemoveNodeByIndex(curExtremumIndex)
		sortedList.AddNodeToHead(nodeExtremum.Value)
	}

	headNode := sortedList.head
	list.head = headNode
}

// SortWayTwo - a little bit faster than SortWayOne, but idea is the same
func (list *LinkedList) SortWayTwo(orderAsc bool) {
	sortedList := LinkedList{} // add nodes to this new list

	for list.head != nil {
		// we need curExtremumPrevNode to remove curExtremumNode
		// without additional traverse in the list
		var curExtremumPrevNode *nodes.Node
		var curExtremumNode *nodes.Node
		curExtremumNode = list.head

		var prev *nodes.Node
		cur := list.head

		for cur != nil {
			if orderAsc {
				// try to get new min/max value
				if cur.Value >= curExtremumNode.Value {
					curExtremumPrevNode = prev
					curExtremumNode = cur
				}
			} else {
				// try to get new min/max value
				if cur.Value <= curExtremumNode.Value {
					curExtremumPrevNode = prev
					curExtremumNode = cur
				}
			}

			// we reached the end of the list, let's remove min/max Node
			if cur.Next == nil {
				if curExtremumPrevNode != nil {
					curExtremumPrevNode.Next = curExtremumNode.Next
				} else {
					list.head = curExtremumNode.Next
				}
				break
			} else {
				prev = cur
				cur = cur.Next
			}
		}
		sortedList.AddNodeToHead(curExtremumNode.Value)
	}

	headNode := sortedList.head
	list.head = headNode
}
