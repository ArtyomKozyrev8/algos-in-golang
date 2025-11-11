package algos_in_golang

import "fmt"

type OrderedAscLinkedList struct {
	head *Node
}

func (list *OrderedAscLinkedList) String() string {
	res := "OAscLL: "
	cur := list.head
	for {
		if cur != nil {
			res += fmt.Sprintf("%d", cur.value)
			cur = cur.next
			if cur != nil {
				res += ","
			}
		} else {
			break
		}
	}

	return res
}

func (list *OrderedAscLinkedList) Append(value int) {
	if list.head == nil {
		list.head = &Node{value: value}
		return
	}

	var prev *Node
	cur := list.head
	for {
		if cur.value <= value {
			prev = cur
			cur = cur.next
			if cur == nil {
				prev.next = &Node{value: value}
				break
			}
		} else {
			if prev == nil {
				list.head = &Node{value: value, next: cur}
			} else {
				prev.next = &Node{value: value, next: cur}
			}
			break
		}
	}
	return
}
