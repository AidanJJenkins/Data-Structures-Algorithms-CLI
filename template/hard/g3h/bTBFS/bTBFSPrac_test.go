package btbfs

import (
	"testing"
)

var tree *BinaryNode = &BinaryNode{
	value: 20,
	right: &BinaryNode{
		value: 50,
		right: &BinaryNode{
			value: 100,
		},
		left: &BinaryNode{
			value: 30,
			right: &BinaryNode{
				value: 45,
			},
			left: &BinaryNode{
				value: 29,
			},
		},
	},
	left: &BinaryNode{
		value: 10,
		right: &BinaryNode{
			value: 15,
		},
		left: &BinaryNode{
			value: 5,
			right: &BinaryNode{
				value: 7,
			},
		},
	},
}

func TestPreOrderSearch(t *testing.T) {
	// Use the "tree" variable from the main package
	first := BFS(tree, 45)
	second := BFS(tree, 45)
	third := BFS(tree, 45)

	if first != true {
		t.Errorf("first should be true, but got %t: ", first)
	}
	if second != true {
		t.Errorf("second should be true, but got %t: ", second)
	}
	if third != true {
		t.Errorf("third should be true, but got %t: ", third)
	}
}
