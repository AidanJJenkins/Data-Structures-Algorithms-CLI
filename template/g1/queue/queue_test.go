package queue

import "testing"

func TestQueue(t *testing.T) {
	list := Queue[int]{0, nil, nil}

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	if value := list.Dequeue(); value != 5 {
		t.Errorf("Expected 5, but got %v", value)
	}
	if list.length != 2 {
		t.Errorf("Expected length 2, but got %v", list.length)
	}

	list.Enqueue(11)

	if value := list.Dequeue(); value != 7 {
		t.Errorf("Expected 7, but got %v", value)
	}
	if value := list.Dequeue(); value != 9 {
		t.Errorf("Expected 9, but got %v", value)
	}
	if value := list.Peek(); value != 11 {
		t.Errorf("Expected 11, but got %v", value)
	}
	if value := list.Dequeue(); value != 11 {
		t.Errorf("Expected 11, but got %v", value)
	}
	if value := list.Dequeue(); value != nil {
		t.Errorf("Expected nil, but got %v", value)
	}
	if list.length != 0 {
		t.Errorf("Expected length 0, but got %v", list.length)
	}

	list.Enqueue(69)
	if value := list.Peek(); value != 69 {
		t.Errorf("Expected 69, but got %v", value)
	}
	if list.length != 1 {
		t.Errorf("Expected length 1, but got %v", list.length)
	}
}
