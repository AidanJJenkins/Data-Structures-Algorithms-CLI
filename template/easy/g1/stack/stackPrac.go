package stack

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	Length int
	Head   *StackNode
}

func NewStack() *Stack {
}

func (s *Stack) Push(value int) {

}

func (s *Stack) Pop() any {

}
func (s *Stack) Peek() any {

}
