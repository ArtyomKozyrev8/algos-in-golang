package main

import (
	"fmt"

	"github.com/ArtyomKozyrev8/algos-in-golang/linkedlist"
	"github.com/ArtyomKozyrev8/algos-in-golang/ordered_linked_list"
)

func main() {
	list := linkedlist.LinkedList{}
	list.AddNodeToTail(3)
	list.AddNodeToTail(4)
	fmt.Println(&list)
	olist := ordered_linked_list.OrderedAscLinkedList{}
	olist.Append(2)
	olist.Append(1)
	olist.Append(0)
	fmt.Println(&olist)
}
