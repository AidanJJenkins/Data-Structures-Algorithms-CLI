package dfsonbst

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

func TestDfs(t *testing.T) {
	// Use the "tree" variable from the main package
	result := DFS(tree, 45)

	if result != true {
		t.Errorf("expected true but got: %t", result)
	}
}
