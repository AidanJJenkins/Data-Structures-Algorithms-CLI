package reversell

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Length     int
	Head, Tail *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func NewSinglyList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

func (ll *LinkedList) Reverse() {

}

func Rec(head *Node) *Node {

}
