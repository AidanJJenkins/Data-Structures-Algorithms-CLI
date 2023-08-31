package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := newSinglyList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if val := list.Get(2); val != 9 {
		t.Errorf("Expected 9 at index 2, but got %v", val)
	}

	if val := list.RemoveAt(1); val != 7 {
		t.Errorf("Expected 7 to be removed, but got %v", val)
	}

	if len := list.length; len != 2 {
		t.Errorf("Expected list length to be 2, but got %v", len)
	}

	list.Append(11)

	if val := list.RemoveAt(1); val != 9 {
		t.Errorf("Expected 9 to be removed, but got %v", val)
	}

	if val := list.Remove(9); val != nil {
		t.Errorf("Expected nil as value when removing non-existing element, but got %v", val)
	}

	if val := list.RemoveAt(0); val != 5 {
		t.Errorf("Expected 5 to be removed, but got %v", val)
	}

	if val := list.RemoveAt(0); val != 11 {
		t.Errorf("Expected 11 to be removed, but got %v", val)
	}

	if len := list.length; len != 0 {
		t.Errorf("Expected list length to be 0, but got %v", len)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if val := list.Get(2); val != 5 {
		t.Errorf("Expected 5 at index 2, but got %v", val)
	}

	if val := list.Get(0); val != 9 {
		t.Errorf("Expected 9 at index 0, but got %v", val)
	}

	if val := list.Remove(9); val != 9 {
		t.Errorf("Expected 9 to be removed, but got %v", val)
	}

	if len := list.length; len != 2 {
		t.Errorf("Expected list length to be 2, but got %v", len)
	}

	if val := list.Get(0); val != 7 {
		t.Errorf("Expected 7 at index 0, but got %v", val)
	}
}

func (ll *LinkedList) PrintList() {
	toPrint := ll.head

	for toPrint != nil {
		fmt.Printf("%d,", toPrint.value.(int))
		toPrint = toPrint.next
	}

	fmt.Printf("\n")
	fmt.Printf("Length: %d\n", ll.length)
}
