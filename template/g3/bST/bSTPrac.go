package bst

type BinaryNode struct {
	value       int
	right, left *BinaryNode
}

type BST struct {
	root *BinaryNode
}

func newBNode(item int) *BinaryNode {
}

func find(node *BinaryNode, item int) bool {
}

func (b *BST) Find(item int) bool {
}

func insert(node *BinaryNode, item int) {
}

func (b *BST) Insert(item int) {
}

func deleteNode(node *BinaryNode, item int) *BinaryNode {
}

func minValueNode(node *BinaryNode) *BinaryNode {
}

func (b *BST) Delete(item int) {
}
