package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	list := NewStack()

	list.Push(5)
	list.Push(7)
	list.Push(9)

	if value := list.Pop(); value != 9 {
		t.Errorf("Expected value 9, but got %v", value)
	}
	if list.length != 2 {
		t.Errorf("Expected length 2, but got %v", list.length)
	}

	list.Push(11)
	if value := list.Pop(); value != 11 {
		t.Errorf("Expected value 11, but got %v", value)
	}
	if value := list.Pop(); value != 7 {
		t.Errorf("Expected value 7, but got %v", value)
	}
	if value := list.Peek(); value != 5 {
		t.Errorf("Expected value 5 (peek), but got %v", value)
	}
	if value := list.Pop(); value != 5 {
		t.Errorf("Expected value 5, but got %v", value)
	}
	if value := list.Pop(); value != nil {
		t.Errorf("Expected nil, but got %v", value)
	}

	list.Push(69)
	if value := list.Peek(); value != 69 {
		t.Errorf("Expected value 69 (peek), but got")
	}
}
