package queue

import "testing"

func TestQueue(t *testing.T) {
	list := new(Queue)

	//testing enqueue
	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	//testing all three methods
	if item, ok := list.Dequeue().(int); ok {
		if item != 5 {
			t.Errorf("Expected %d, received %d", 5, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if list.length != 2 {
		t.Errorf("expected %d recieved %d", 2, list.length)
	}

	list.Enqueue(11)

	if item, ok := list.Dequeue().(int); ok {
		if item != 7 {
			t.Errorf("Expected %d, received %d", 7, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if item, ok := list.Dequeue().(int); ok {
		if item != 9 {
			t.Errorf("Expected %d, received %d", 9, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if item, ok := list.Peek().(int); ok {
		if item != 11 {
			t.Errorf("Expected %d, received %d", 11, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if item, ok := list.Dequeue().(int); ok {
		if item != 11 {
			t.Errorf("Expected %d, received %d", 11, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if item, ok := list.Dequeue().(int); ok {
		if item != -1 {
			t.Errorf("Expected %d, received %d", -1, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if list.length != 0 {
		t.Errorf("expected %d recieved %d", 0, list.length)
	}

	//making sure it works when theres no more list
	list.Enqueue(69)

	if item, ok := list.Peek().(int); ok {
		if item != 69 {
			t.Errorf("Expected %d, received %d", 69, item)
		}
	} else {
		t.Error("Unexpected value returned from Dequeue")
	}

	if list.length != 1 {
		t.Errorf("expected %d recieved %d", 1, list.length)
	}
}
