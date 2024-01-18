package btpostorder

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
	result := PostOrderSearch(tree)

	// Define the expected output
	expected := []int{7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20}

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderSearch(tree) = %v; want %v", result, expected)
	}
}
