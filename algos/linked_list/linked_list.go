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
	if len(vals) == 0 {
		var head *ListNode
		return head
	}

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

func MergeTwoSortedListAsSortedList(list1 *ListNode, list2 *ListNode) *ListNode {
	var lastNode *ListNode  // we work with the node to attach new nodes
	var firstNode *ListNode //  the first node of new list - used only once - then returned in the end of func
	cur1 := list1
	cur2 := list2
	for {
		if firstNode != nil && lastNode == nil {
			lastNode = firstNode
		}

		if cur1 == nil && cur2 == nil {
			break
		} else if cur1 == nil {
			if firstNode == nil {
				firstNode = cur2
			} else {
				lastNode.Next = cur2
				lastNode = cur2
			}
			cur2 = cur2.Next
		} else if cur2 == nil {
			if firstNode == nil {
				firstNode = cur1
			} else {
				lastNode.Next = cur1
				lastNode = cur1
			}
			cur1 = cur1.Next
		} else {
			if cur1.Val <= cur2.Val {
				if firstNode == nil {
					firstNode = cur1
				} else {
					lastNode.Next = cur1
					lastNode = cur1
				}
				cur1 = cur1.Next
			} else {
				if firstNode == nil {
					firstNode = cur2
				} else {
					lastNode.Next = cur2
					lastNode = cur2
				}
				cur2 = cur2.Next
			}
		}
	}

	if lastNode != nil { // prevent Goland from complaining about possible null reference
		lastNode.Next = nil // actually we do not need to do it, because last node in both lists have Next = nil
	}
	return firstNode
}

func FindMiddleNodeValue(head *ListNode) int {
	if head == nil {
		return -1
	}

	curFast := head // do two steps
	curSlow := head // do one step

	for {
		//try to make two steps
		curFast = curFast.Next
		if curFast == nil { // check results of first step
			return curSlow.Val
		}
		curFast = curFast.Next
		if curFast == nil { // check results of second step
			return curSlow.Next.Val
		}
		curSlow = curSlow.Next // make only one step
	}
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	curFast := head
	curSlow := head
	prevCurSlow := new(ListNode)
	steps := n

	for {
		curFast = curFast.Next
		if steps == 0 {
			prevCurSlow = curSlow
			curSlow = curSlow.Next
		} else {
			steps--
		}

		if curFast == nil {
			if steps > 0 {
				return head.Next // we remove first (don't know why!!!)
			} else {
				if curSlow == head {
					head = head.Next
				} else {
					if curSlow.Next != nil {
						prevCurSlow.Next = curSlow.Next // detach curSlow node from list
					} else {
						prevCurSlow.Next = nil
					}
				}
				break
			}
		}
	}

	return head
}
