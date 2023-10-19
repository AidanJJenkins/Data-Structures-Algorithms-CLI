package bst

import "testing"

func TestBST(t *testing.T) {
	// Use the "tree" variable from the main package
	tree := &BST{nil}

	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(9)
	tree.Insert(11)
	seven := tree.Find(7)
	tree.Delete(11)
	eleven := tree.Find(11)

	if seven != true {
		t.Errorf("expected true, got: %t", seven)
	}
	if eleven != false {
		t.Errorf("expected false, got: %t", eleven)
	}
}
