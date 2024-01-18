package btpreorder

import (
	"reflect"
	"testing"
)

var tree *BinaryNode = &BinaryNode{
	Val: 20,
	Right: &BinaryNode{
		Val: 50,
		Right: &BinaryNode{
			Val: 100,
		},
		Left: &BinaryNode{
			Val: 30,
			Right: &BinaryNode{
				Val: 45,
			},
			Left: &BinaryNode{
				Val: 29,
			},
		},
	},
	Left: &BinaryNode{
		Val: 10,
		Right: &BinaryNode{
			Val: 15,
		},
		Left: &BinaryNode{
			Val: 5,
			Right: &BinaryNode{
				Val: 7,
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
