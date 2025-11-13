package main

import (
	"fmt"

	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/linkedlist"
	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/ordered_linked_list"
)

func main() {
	list := linkedlist.LinkedList{}
	list.AddNodeToTail(3)
	list.AddNodeToTail(4)
	fmt.Println(&list)

	oList := ordered_linked_list.OrderedAscLinkedList{}
	oList.Append(2)
	oList.Append(1)
	oList.Append(0)
	fmt.Println(&oList)
}
