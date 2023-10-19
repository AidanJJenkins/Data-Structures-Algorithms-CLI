package stack

import (
	"testing"
)

func TestStackPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	out := s.Pop()
	if out != 3 {
		t.Errorf("Expected Pop() to return 3, got %d", out)
	}
	if s.length != 2 {
		t.Errorf("Expected length to be 2 after Pop, got %d", s.length)
	}

	out = s.Pop()
	if out != 2 {
		t.Errorf("Expected Pop() to return 2, got %d", out)
	}
	if s.length != 1 {
		t.Errorf("Expected length to be 1 after Pop, got %d", s.length)
	}

	s = NewStack[int]()
	out = s.Pop()
	if out != nil {
		t.Errorf("Expected Pop() on an empty stack to return nil, got %d", out)
	}
	if s.length != 0 {
		t.Errorf("Expected length of an empty stack to be 0, got %d", s.length)
	}

	t.Logf("TestStackPop passed")
}
