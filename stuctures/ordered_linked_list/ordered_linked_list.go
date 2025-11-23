package ordered_linked_list

import (
	"fmt"

	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/nodes"
)

type OLLError struct {
	value string
}

func (e *OLLError) Error() string {
	return fmt.Sprintf("OLLError: %v", e.value)
}

type OrderedAscLinkedList struct {
	head *nodes.Node
}

func (list *OrderedAscLinkedList) String() string {
	res := "OAscLL: "
	cur := list.head
	for {
		if cur != nil {
			res += fmt.Sprintf("%d", cur.Value)
			cur = cur.Next
			if cur != nil {
				res += ","
			}
		} else {
			break
		}
	}

	return res
}

func (list *OrderedAscLinkedList) RemoveByValue(value int) (int, error) {
	var prev *nodes.Node
	cur := list.head

	for cur != nil {
		if cur.Value == value {
			if prev != nil {
				prev.Next = cur.Next
			} else {
				list.head = cur.Next
			}
			return cur.Value, nil
		} else {
			if cur.Value > value {
				break // our list is ordered, if cur.value > value, we can stop searching
			}
			prev = cur
			cur = cur.Next
		}
	}

	return 0, &OLLError{fmt.Sprintf("value %d not found", value)}
}

func (list *OrderedAscLinkedList) Append(value int) {
	if list.head == nil {
		list.head = &nodes.Node{Value: value}
		return
	}

	var prev *nodes.Node
	cur := list.head
	for {
		if cur.Value <= value {
			prev = cur
			cur = cur.Next
			if cur == nil {
				prev.Next = &nodes.Node{Value: value}
				break
			}
		} else {
			if prev == nil {
				list.head = &nodes.Node{Value: value, Next: cur}
			} else {
				prev.Next = &nodes.Node{Value: value, Next: cur}
			}
			break
		}
	}
	return
}

func (list *OrderedAscLinkedList) RemoveByIndex(index int) (int, error) {
	var curIndex int = 0
	var prev *nodes.Node
	cur := list.head

	for cur != nil {
		if curIndex == index {
			if prev == nil {
				list.head = cur.Next
			} else {
				prev.Next = cur.Next
			}
			return cur.Value, nil
		} else {
			prev = cur
			cur = cur.Next
			curIndex = curIndex + 1
		}
	}

	return 0, &OLLError{fmt.Sprintf("index %d out of range", index)}
}

func (list *OrderedAscLinkedList) CheckValueInList(value int) bool {
	cur := list.head
	for cur != nil {
		if cur.Value == value {
			return true
		}
		cur = cur.Next
	}
	return false
}
