package btbfs

import (
	"testing"
)

var tree *TreeNode = &TreeNode{
	Val: 20,
	Right: &TreeNode{
		Val: 50,
		Right: &TreeNode{
			Val: 100,
		},
		Left: &TreeNode{
			Val: 30,
			Right: &TreeNode{
				Val: 45,
			},
			Left: &TreeNode{
				Val: 29,
			},
		},
	},
	Left: &TreeNode{
		Val: 10,
		Right: &TreeNode{
			Val: 15,
		},
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 7,
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
