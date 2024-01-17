package invertbtree

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

var invertedTree *TreeNode = &TreeNode{
	Val: 20,
	Left: &TreeNode{
		Val: 50,
		Left: &TreeNode{
			Val: 100,
		},
		Right: &TreeNode{
			Val: 30,
			Left: &TreeNode{
				Val: 45,
			},
			Right: &TreeNode{
				Val: 29,
			},
		},
	},
	Right: &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 15,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 7,
			},
		},
	},
}

var badTree *TreeNode = &TreeNode{
	Val: 20,
	Left: &TreeNode{
		Val: 50,
		Left: &TreeNode{
			Val: 100,
		},
		Right: &TreeNode{
			Val: 30,
			Left: &TreeNode{
				Val: 45,
			},
			Right: &TreeNode{
				Val: 29,
			},
		},
	},
}

func TestInvert(t *testing.T) {
	testTree := Invert(tree)
	result := Compare(testTree, invertedTree)

	if result != true {
		t.Errorf("expected true but got: %t", result)
	}

	testTree2 := Invert(badTree)
	result2 := Compare(testTree2, invertedTree)

	if result2 != false {
		t.Errorf("expected true but got: %t", result)
	}

}

func Compare(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
