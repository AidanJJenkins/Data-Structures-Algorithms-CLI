package btpreorder

import (
	"reflect"
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
	result := PreOrderSearch(tree)

	// Define the expected output
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderSearch(tree) = %v; want %v", result, expected)
	}
}
