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

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
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

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	currencyCode["Pip"] = "OIOPIO"
	currency := "USD"
	currencyName := currencyCode[currency]
	fmt.Println("Currency name for currency code", currency, "is", currencyName)

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}
