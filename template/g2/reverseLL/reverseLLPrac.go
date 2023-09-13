package reversell

type value any

type Node struct {
	value value
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func newNode(value value) *Node {
	return &Node{value, nil}
}

func newSinglyList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

func (ll *LinkedList) Reverse() {

}

func Rec(head *Node) *Node {

}
