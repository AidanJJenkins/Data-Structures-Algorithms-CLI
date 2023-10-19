package doublylinkedlist

type Node[T comparable] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	Length     int
	head, tail *Node[T]
}

func newNode[T comparable](value T) *Node[T] {
}

func newDoublyList[T comparable]() *DoublyLinkedList[T] {
}

func (ll *DoublyLinkedList[T]) Prepend(item T) {

}

func (ll *DoublyLinkedList[T]) Append(item T) {

}

func (ll *DoublyLinkedList[T]) InsertAt(item T, idx int) {

}

func (ll *DoublyLinkedList[T]) Remove(item any) any {

}

func (ll *DoublyLinkedList[T]) removeNode(node *Node[T]) T {

}

func (ll *DoublyLinkedList[T]) Get(idx int) any {

}

func (ll *DoublyLinkedList[T]) RemoveAt(idx int) any {

}

func (ll *DoublyLinkedList[T]) getAt(idx int) *Node[T] {

}
