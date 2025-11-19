package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) AddToTail(vals ...int) {
	cur := ln
	var lastNode *ListNode
	for {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			lastNode = &ListNode{Val: vals[0], Next: nil}
			cur.Next = lastNode
			break
		}
	}

	for i, val := range vals {
		if i > 0 {
			newLastNode := &ListNode{Val: val, Next: nil}
			lastNode.Next = newLastNode
			lastNode = newLastNode
		}
	}
}

func (ln *ListNode) PrintList() string {
	if ln == nil {
		return "[]"
	}

	resStr := ""
	for cur := ln; cur != nil; cur = cur.Next {
		if resStr == "" {
			resStr += fmt.Sprintf("%v", cur.Val)
		} else {
			resStr += fmt.Sprintf(",%v", cur.Val)
		}
	}
	return fmt.Sprintf("[%v]", resStr)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	rest := 0

	var firstNodeOfNewList *ListNode
	var curNodeOfNewList *ListNode
	for {
		if cur1 == nil && cur2 == nil && rest == 0 {
			break
		}

		sumRes := 0
		if cur1 != nil && cur2 != nil {
			sumRes = cur1.Val + cur2.Val + rest
			cur1 = cur1.Next
			cur2 = cur2.Next
		} else if cur1 != nil {
			sumRes = cur1.Val + rest
			cur1 = cur1.Next
		} else if cur2 != nil {
			sumRes = cur2.Val + rest
			cur2 = cur2.Next
		} else {
			sumRes = rest
		}

		nodeVal := 0
		if sumRes <= 9 {
			nodeVal = sumRes
			rest = 0
		} else {
			nodeVal = sumRes - 10
			rest = 1
		}

		if firstNodeOfNewList == nil {
			firstNodeOfNewList = &ListNode{nodeVal, nil}
			curNodeOfNewList = firstNodeOfNewList
		} else {
			if curNodeOfNewList != nil {
				curNodeOfNewList.Next = &ListNode{nodeVal, nil}
				curNodeOfNewList = curNodeOfNewList.Next
			}
		}

	}

	return firstNodeOfNewList
}
