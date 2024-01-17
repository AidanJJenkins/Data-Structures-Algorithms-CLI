package comparetrees

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

var tree2 *TreeNode = &TreeNode{
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
				Val: 45,
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
