package algos_in_golang

import "strconv"

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

// PrintList converts nodes in LinkedList to string
func (list *LinkedList) PrintList() string {
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
