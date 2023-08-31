package comparetrees

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

var tree2 *BinaryNode = &BinaryNode{
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
				value: 45,
			},
		},
	},
}

func TestCompare(t *testing.T) {
	comp := Compare(tree, tree2)
	comp2 := Compare(tree, tree)
	comp3 := Compare(tree2, tree2)
	if comp != false {
		t.Errorf("expected false, got: %t", comp)
	}
	if comp2 != true {
		t.Errorf("expected true, got: %t", comp2)
	}
	if comp3 != true {
		t.Errorf("expected true, got: %t", comp3)
	}
}
