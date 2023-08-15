package doublylinkedlist

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	Length int
	head   *Node
	tail   *Node
}

func newNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func newDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (ll *DoublyLinkedList) Prepend(item int) {
}

func (ll *DoublyLinkedList) InsertAt(item int, idx int) {
}

func (ll *DoublyLinkedList) Append(item int) {
}

func (ll *DoublyLinkedList) Remove(item int) int {
}

func (ll *DoublyLinkedList) removeNode(node *Node) int {
}

func (ll *DoublyLinkedList) Get(idx int) int {
}

func (ll *DoublyLinkedList) RemoveAt(idx int) int {
}

func (ll *DoublyLinkedList) getAt(idx int) *Node {
}
