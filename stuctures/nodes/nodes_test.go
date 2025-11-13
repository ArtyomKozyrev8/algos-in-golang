package nodes

import (
	"fmt"
	"testing"
)

func TestNode_String(t *testing.T) {
	node := &Node{Value: 10}

	if node.String() != "Node(10)" {
		t.Error("Expected Node(10) got ", node.String())
	}

	node.Next = &Node{Value: 20}

	if node.String() != "Node(10)" {
		t.Error("Expected Node(10) got ", node.String())
	}

	if node.Next.String() != "Node(20)" {
		t.Error("Expected Node(20) got ", node.Next.String())
	}

	if fmt.Sprintf("%v", node) != "Node(10)" {
		t.Error("Expected Node(10) got ", fmt.Sprintf("%v", node))
	}
}

func TestNode_IsEqual(t *testing.T) {
	node := &Node{Value: 10}
	anotherOne := &Node{Value: 10}
	if !node.IsEqual(anotherOne) {
		t.Error("Expected Node(10) == Node(10)")
	}

	anotherTwo := &Node{Value: 20}
	if node.IsEqual(anotherTwo) {
		t.Error("Expected Node(10) != Node(20)")
	}
}

func TestNode_IsGreater(t *testing.T) {
	node := &Node{Value: 10}
	anotherOne := &Node{Value: 20}

	if node.IsGreater(anotherOne) {
		t.Error("Expected Node(10) < Node(20)")
	}

	anotherTwo := &Node{Value: 5}
	if !node.IsGreater(anotherTwo) {
		t.Error("Expected Node(10) > Node(5)")
	}
}

func TestNode_IsGreaterOrEqual(t *testing.T) {
	node := &Node{Value: 10}
	anotherOne := &Node{Value: 20}

	if node.IsGreaterOrEqual(anotherOne) {
		t.Error("Expected Node(10) < Node(20)")
	}

	anotherTwo := &Node{Value: 10}
	if !node.IsGreaterOrEqual(anotherTwo) {
		t.Error("Expected Node(10) >= Node(10)")
	}

	anotherThree := &Node{Value: 5}
	if !node.IsGreaterOrEqual(anotherThree) {
		t.Error("Expected Node(10) >= Node(5)")
	}
}
