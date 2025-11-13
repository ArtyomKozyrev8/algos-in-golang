package nodes

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("Node(%d)", node.Value)
}

func (node *Node) IsGreaterOrEqual(another *Node) bool {
	return node.Value >= another.Value
}

func (node *Node) IsEqual(another *Node) bool {
	return node.Value == another.Value
}

func (node *Node) IsGreater(another *Node) bool {
	return node.Value > another.Value
}
