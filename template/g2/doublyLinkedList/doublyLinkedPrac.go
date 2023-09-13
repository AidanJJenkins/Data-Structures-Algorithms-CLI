package doublylinkedlist

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	Length int
	head   *Node
	tail   *Node
}

func newNode(value int) *Node {
}

func newDoublyList() *DoublyLinkedList {
}

func (ll *DoublyLinkedList) Prepend(item any) {
}

func (ll *DoublyLinkedList) Append(item int) {
}

func (ll *DoublyLinkedList) InsertAt(item any, idx int) {
}

func (ll *DoublyLinkedList) Remove(item any) any {
}

func (ll *DoublyLinkedList) removeNode(node *Node) any {
}

func (ll *DoublyLinkedList) Get(idx int) any {
}

func (ll *DoublyLinkedList) RemoveAt(idx int) any {
}

func (ll *DoublyLinkedList) getAt(idx int) *Node {
}
