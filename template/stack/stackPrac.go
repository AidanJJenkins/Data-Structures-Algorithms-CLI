package stack

type Node struct {
	value any
	prev  *Node
}

type Stack struct {
	length int
	head   *Node
}

func (s *Stack) push(item any) {
}

func (s *Stack) pop() any {
}
func (s *Stack) peak() any {
}
