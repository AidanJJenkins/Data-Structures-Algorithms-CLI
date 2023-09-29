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

func TestEnqueue(t *testing.T) {
	list := Queue[int]{0, nil, nil}

	list.Enqueue(1)
	if list.length != 1 {
		t.Errorf("Expected length to be 1, got %d", list.length)
	}
	if list.head.value != 1 {
		t.Errorf("Expected head value to be 1, got %d", list.head.value)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := Queue[int]{0, nil, nil}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	out := q.Dequeue()
	if out != 1 {
		t.Errorf("Expected Dequeue() to return 1, got %d", out)
	}
	if q.length != 2 {
		t.Errorf("Expected length to be 2 after Dequeue, got %d", q.length)
	}

	out = q.Dequeue()
	if out != 2 {
		t.Errorf("Expected Dequeue() to return 2, got %d", out)
	}
	if q.length != 1 {
		t.Errorf("Expected length to be 1 after Dequeue, got %d", q.length)
	}

	newQ := NewQueue[int]()
	out = newQ.Dequeue()
	if out != nil {
		t.Errorf("Expected Dequeue() on an empty queue to return nil, got %d", out)
	}
	if newQ.length != 0 {
		t.Errorf("Expected length of an empty queue to be 0, got %d", q.length)
	}
}
