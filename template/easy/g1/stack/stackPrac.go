package stack

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Stack[T comparable] struct {
	length int
	head   *Node[T]
}

func NewStack[T comparable]() *Stack[T] {
}

func (s *Stack[T]) Push(item T) {

}

func (s *Stack[T]) Pop() any {

}
func (s *Stack[T]) Peek() any {

}
