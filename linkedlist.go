package algos_in_golang

import (
	"errors"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

// AddNodeToTail adds new node to LinkedList
func (list *LinkedList) AddNodeToTail(value int) {
	if list.head == nil {
		list.head = &Node{value, nil}
	} else {
		cur := list.head
		for {
			if cur.next == nil {
				cur.next = &Node{value, nil}
				break
			}
			cur = cur.next
		}
	}
}

// String converts nodes in LinkedList to string
func (list *LinkedList) String() string {
	res := "LinkedList: "
	cur := list.head

	for cur != nil {
		res += strconv.Itoa(cur.value)
		if cur.next != nil {
			res += ","
		}
		cur = cur.next
	}

	return res
}

// ReverseList reverse order of nodes in LinkedList
func (list *LinkedList) ReverseList() {
	var prev *Node
	cur := list.head

	for cur != nil {
		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}

	list.head = prev
}

// GetNodeByIndex returns Node from LinkedList by index
func (list *LinkedList) GetNodeByIndex(index int) *Node {
	cur := list.head
	curIndex := 0
	for cur != nil {
		if curIndex == index {
			return cur
		}
		cur = cur.next
		curIndex += 1
	}
	return nil
}

// GetNodeByValue returns Node from LinkedList by value
func (list *LinkedList) GetNodeByValue(value int) *Node {
	cur := list.head
	for cur != nil {
		if cur.value == value {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// AddNodeToHead adds new node to the head of LinkedList
func (list *LinkedList) AddNodeToHead(value int) {
	if list.head == nil {
		list.head = &Node{value, nil}
	} else {
		cur := list.head
		list.head = &Node{value, cur}
	}
}

// RemoveNodeFromHead removes and returns Node from head of LinkedList
func (list *LinkedList) RemoveNodeFromHead() *Node {
	if list.head == nil {
		return nil
	}

	cur := list.head
	curNext := cur.next
	list.head = curNext

	return cur
}

// RemoveNodeFromTail removes and returns Node from tail of LinkedList
func (list *LinkedList) RemoveNodeFromTail() *Node {
	if list.head == nil {
		return nil
	}

	var prev *Node
	cur := list.head
	for {
		if cur.next != nil {
			prev = cur
			cur = cur.next
		} else {
			if prev == nil {
				list.head = nil
			} else {
				prev.next = nil
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
		cur = cur.next
		num += 1
	}

	return num
}

// RemoveNodeByIndex removes Node from LinkedList based on index
func (list *LinkedList) RemoveNodeByIndex(index int) (error, *Node) {
	if index < 0 {
		return errors.New("index can't be negative number"), nil
	}

	if list.head == nil {
		return errors.New("can't remove from empty list"), nil
	}

	curIndex := 0
	var prev *Node
	cur := list.head

	for cur != nil {
		if curIndex == index {
			if prev == nil {
				list.head = cur.next
			} else {
				prev.next = cur.next
			}
			return nil, cur
		} else {
			prev = cur
			cur = cur.next
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
		curExtremumValue := list.head.value
		for cur != nil {
			if cur.next == nil {
				break
			} else {
				curIndex += 1
				if orderAsc == true {
					if curExtremumValue <= cur.next.value {
						curExtremumValue = cur.next.value
						curExtremumIndex = curIndex
					}
				} else {
					if curExtremumValue >= cur.next.value {
						curExtremumValue = cur.next.value
						curExtremumIndex = curIndex
					}
				}
				cur = cur.next
			}
		}
		_, nodeExtremum := list.RemoveNodeByIndex(curExtremumIndex)
		sortedList.AddNodeToHead(nodeExtremum.value)
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
		var curExtremumPrevNode *Node
		var curExtremumNode *Node
		curExtremumNode = list.head

		var prev *Node
		cur := list.head

		for cur != nil {
			if orderAsc {
				// try to get new min/max value
				if cur.value >= curExtremumNode.value {
					curExtremumPrevNode = prev
					curExtremumNode = cur
				}
			} else {
				// try to get new min/max value
				if cur.value <= curExtremumNode.value {
					curExtremumPrevNode = prev
					curExtremumNode = cur
				}
			}

			// we reached the end of the list, let's remove min/max Node
			if cur.next == nil {
				if curExtremumPrevNode != nil {
					curExtremumPrevNode.next = curExtremumNode.next
				} else {
					list.head = curExtremumNode.next
				}
				break
			} else {
				prev = cur
				cur = cur.next
			}
		}
		sortedList.AddNodeToHead(curExtremumNode.value)
	}

	headNode := sortedList.head
	list.head = headNode
}
