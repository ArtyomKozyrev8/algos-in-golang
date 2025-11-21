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

func BuildList(vals ...int) *ListNode {
	firstNode := &ListNode{Val: vals[0]}
	if len(vals) > 1 {
		firstNode.AddToTail(vals[1:]...)
	}
	return firstNode
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

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	firstNode := head

	totalNodeNumber := 0
	for tempK := k; tempK > 0; tempK-- {
		var prev *ListNode
		cur := firstNode

		for cur.Next != nil {
			prev = cur
			cur = cur.Next
			totalNodeNumber += 1
		}

		totalNodeNumber += 1 // // we add one because previous loop was cur.Next != nil

		if k > totalNodeNumber {
			break
		}

		if prev != nil {
			prev.Next = nil
		}
		cur.Next = firstNode
		firstNode = cur
	}

	// k > totalNodeNumber means that list do full rotation and is the same
	if k > totalNodeNumber {
		// we need to check only additional steps
		newK := k % totalNodeNumber
		firstNode = head
		for ; newK > 0; newK-- {
			var prev *ListNode
			cur := firstNode

			for cur.Next != nil {
				prev = cur
				cur = cur.Next
			}

			if prev != nil {
				prev.Next = nil
			}
			cur.Next = firstNode
			firstNode = cur
		}
	}

	return firstNode
}
