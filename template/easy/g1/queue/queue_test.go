package queue

import "testing"

func TestQueue(t *testing.T) {
	list := Queue{0, nil, nil}

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	if value := list.Dequeue(); value != 5 {
		t.Errorf("Expected 5, but got %v", value)
	}
	if list.Length != 2 {
		t.Errorf("Expected Length 2, but got %v", list.Length)
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
	if list.Length != 0 {
		t.Errorf("Expected Length 0, but got %v", list.Length)
	}

	list.Enqueue(69)
	if value := list.Peek(); value != 69 {
		t.Errorf("Expected 69, but got %v", value)
	}
	if list.Length != 1 {
		t.Errorf("Expected Length 1, but got %v", list.Length)
	}
}

func TestEnqueue(t *testing.T) {
	list := Queue{0, nil, nil}

	list.Enqueue(1)
	if list.Length != 1 {
		t.Errorf("Expected Length to be 1, got %d", list.Length)
	}
	if list.Head.Val != 1 {
		t.Errorf("Expected head value to be 1, got %d", list.Head.Val)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := Queue{0, nil, nil}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	out := q.Dequeue()
	if out != 1 {
		t.Errorf("Expected Dequeue() to return 1, got %d", out)
	}
	if q.Length != 2 {
		t.Errorf("Expected Length to be 2 after Dequeue, got %d", q.Length)
	}

	out = q.Dequeue()
	if out != 2 {
		t.Errorf("Expected Dequeue() to return 2, got %d", out)
	}
	if q.Length != 1 {
		t.Errorf("Expected Length to be 1 after Dequeue, got %d", q.Length)
	}

	newQ := NewQueue()
	out = newQ.Dequeue()
	if out != nil {
		t.Errorf("Expected Dequeue() on an empty queue to return nil, got %d", out)
	}
	if newQ.Length != 0 {
		t.Errorf("Expected Length of an empty queue to be 0, got %d", q.Length)
	}
}
