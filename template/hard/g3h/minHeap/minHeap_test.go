package minheap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	if heap.Length != 0 {
		t.Errorf("Expected length to be 0, but got %d", heap.Length)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(70)
	heap.Insert(500)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if heap.Length != 8 {
		t.Errorf("Expected length to be 8, but got %d", heap.Length)
	}

	if heap.Delete() != 1 {
		t.Errorf("Expected deleted value to be 1, but got %d", heap.Delete())
	}
	if heap.Delete() != 3 {
		t.Errorf("Expected deleted value to be 3, but got %d", heap.Delete())
	}
	if heap.Delete() != 4 {
		t.Errorf("Expected deleted value to be 4, but got %d", heap.Delete())
	}
	if heap.Delete() != 5 {
		t.Errorf("Expected deleted value to be 5, but got %d", heap.Delete())
	}

	if heap.Length != 4 {
		t.Errorf("Expected length to be 4, but got %d", heap.Length)
	}

	if heap.Delete() != 7 {
		t.Errorf("Expected deleted value to be 7, but got %d", heap.Delete())
	}
	if heap.Delete() != 8 {
		t.Errorf("Expected deleted value to be 8, but got %d", heap.Delete())
	}
	if heap.Delete() != 70 {
		t.Errorf("Expected deleted value to be 70, but got %d", heap.Delete())
	}
	if heap.Delete() != 500 {
		t.Errorf("Expected deleted value to be 500, but got %d", heap.Delete())
	}

	if heap.Length != 0 {
		t.Errorf("Expected length to be 0, but got %d", heap.Length)
	}
}
