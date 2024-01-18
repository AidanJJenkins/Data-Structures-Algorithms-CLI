package btinorder

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
	result := InOrderSearch(tree)

	// Define the expected output
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderSearch(tree) = %v; want %v", result, expected)
	}
}
