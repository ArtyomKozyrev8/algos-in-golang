package main

import (
	"fmt"

	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/linkedlist"
	"github.com/ArtyomKozyrev8/algos-in-golang/stuctures/ordered_linked_list"
)

func Smth(n int8) {
	for i := n; i >= 0; i-- {
		str := ""
		for j := i; j > 0; j-- {
			fmt.Print(str, "*")
		}
		fmt.Println()
	}
}

func Fun() {
label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i+j >= 15 {
				fmt.Println(fmt.Sprintf("Exit i=%v | j=%v", i, j))
				break label1
			} else {
				fmt.Println(fmt.Sprintf("i=%v | j=%v", i, j))
			}
		}
	}
}

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

	Smth(8)

	Fun()
}
