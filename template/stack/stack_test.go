package stack

import "testing"

func TestStack(t *testing.T) {
	stack := &Stack{}

	// Push elements into the stack
	stack.push(5)
	stack.push(7)
	stack.push(9)

	// test all methods together
	if popped := stack.pop(); popped != 9 {
		t.Errorf("Expected 9, but got %v", popped)
	}

	if stack.length != 2 {
		t.Errorf("Expected 2, but got %v", stack.length)
	}

	stack.push(11)
	if popped := stack.pop(); popped != 11 {
		t.Errorf("Expected 11, but got %v", popped)
	}

	if popped := stack.pop(); popped != 7 {
		t.Errorf("Expected 7, but got %v", popped)
	}

	if peak := stack.peak(); peak != 5 {
		t.Errorf("Expected 5, but got %v", peak)
	}

	if popped := stack.pop(); popped != 5 {
		t.Errorf("Expected 5, but got %v", popped)
	}

	// Check if the stack is empty after popping all elements
	if stack.length != 0 || stack.head != nil {
		t.Errorf("Stack should be empty, but it has length %d", stack.length)
	}

	// Try popping from an empty stack
	if popped := stack.pop(); popped != nil {
		t.Errorf("Expected nil from an empty stack, but got %v", popped)
	}

	stack.push(69)
	if peak := stack.peak(); peak != 69 {
		t.Errorf("Expected 69, but got %v", peak)
	}

	if stack.length != 1 {
		t.Errorf("Expected 1, but got %v", stack.length)
	}
}
