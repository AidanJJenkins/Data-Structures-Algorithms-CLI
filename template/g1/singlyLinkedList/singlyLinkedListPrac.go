package singlylinkedlist

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	length     int
	head, tail *Node[T]
}

func newNode[T comparable](value T) *Node[T] {
}

func newSinglyList[T comparable]() *LinkedList[T] {
}

func (ll *LinkedList[T]) Prepend(item T) {

}

func (ll *LinkedList[T]) Append(item T) {

}

func (ll *LinkedList[T]) Remove(item T) any {

}

func (ll *LinkedList[T]) Get(idx int) any {

}

func (ll *LinkedList[T]) InsertAt(item T, idx int) {

}

func (ll *LinkedList[T]) RemoveAt(idx int) any {

}
