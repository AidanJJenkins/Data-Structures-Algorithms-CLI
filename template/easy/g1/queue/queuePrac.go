package queue

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Queue[T comparable] struct {
	length     int
	head, tail *Node[T]
}

func NewQueue[T comparable]() *Queue[T] {
}

func (q *Queue[T]) Enqueue(item T) {

}

func (q *Queue[T]) Dequeue() any {

}

func (q *Queue[T]) Peek() any {
}
