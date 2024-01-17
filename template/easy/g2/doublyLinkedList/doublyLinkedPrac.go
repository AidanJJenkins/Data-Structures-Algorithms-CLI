package doublylinkedlist

type Node struct {
	Val        int
	Next, Prev *Node
}

type DoublyLinkedList struct {
	Length     int
	Head, Tail *Node
}

func NewNode(value int) *Node {
}

func NewDoublyList() *DoublyLinkedList {
}

func (ll *DoublyLinkedList) Prepend(value int) {

}

func (ll *DoublyLinkedList) Append(value int) {

}

func (ll *DoublyLinkedList) removeNode(node *Node) int {

}

func (ll *DoublyLinkedList) getAt(idx int) *Node {

}

func (ll *DoublyLinkedList) InsertAt(value int, idx int) {

}

func (ll *DoublyLinkedList) Remove(value int) any {

}

func (ll *DoublyLinkedList) Get(idx int) any {

}

func (ll *DoublyLinkedList) RemoveAt(idx int) any {

}
