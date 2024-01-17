package singlylinkedlist

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Length     int
	Head, Tail *Node
}

func NewNode(value int) *Node {
}

func NewSinglyList() *LinkedList {
}

func (ll *LinkedList) Prepend(value int) {

}

func (ll *LinkedList) Append(value int) {

}

func (ll *LinkedList) Remove(value int) any {

}

func (ll *LinkedList) Get(idx int) any {

}

func (ll *LinkedList) InsertAt(value int, idx int) {

}

func (ll *LinkedList) RemoveAt(idx int) any {

}
