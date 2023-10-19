package reversell

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	length     int
	head, tail *Node[T]
}

func newNode[T comparable](value T) *Node[T] {
	return &Node[T]{value, nil}
}

func newSinglyList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{0, nil, nil}
}

func (ll *LinkedList[T]) Reverse() {

}

func Rec[T comparable](head *Node[T]) *Node[T] {
}
