package stack

type val = any

type Node struct {
	value val
	next  *Node
}

type Stack struct {
	length int
	head   *Node
}

func NewStack() *Stack {
}

func (s *Stack) Push(item val) {
}

func (s *Stack) Pop() any {
}
func (s *Stack) Peek() any {
}
