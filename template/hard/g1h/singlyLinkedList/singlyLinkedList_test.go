package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := NewSinglyList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if val := list.Get(2); val != 9 {
		t.Errorf("Expected 9 at index 2, but got %v", val)
	}

	if val := list.RemoveAt(1); val != 7 {
		t.Errorf("Expected 7 to be removed, but got %v", val)
	}

	if len := list.Length; len != 2 {
		t.Errorf("Expected list Length to be 2, but got %v", len)
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

	if len := list.Length; len != 0 {
		t.Errorf("Expected list Length to be 0, but got %v", len)
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

	if len := list.Length; len != 2 {
		t.Errorf("Expected list Length to be 2, but got %v", len)
	}

	if val := list.Get(0); val != 7 {
		t.Errorf("Expected 7 at index 0, but got %v", val)
	}
}

func (ll *LinkedList) PrintList() {
	toPrint := ll.Head

	for toPrint != nil {
		fmt.Printf("%v,", toPrint.Val)
		toPrint = toPrint.Next
	}

	fmt.Printf("\n")
	fmt.Printf("Length: %d\n", ll.Length)
}

func EdgeTest(t *testing.T) {
	list := NewSinglyList()

	list.Append(5)
	if val := list.Remove(5); val != 5 {
		t.Errorf("Expected 5, but got %v", val)
	}

	if val := list.Get(4); val != nil {
		t.Errorf("Expected nil at index 4, but got %v", val)
	}

	list.Prepend(6)
	list.Prepend(10)

	if val := list.RemoveAt(4); val != nil {
		t.Errorf("Expected nil at index 4, but got %v", val)
	}
	if val := list.RemoveAt(0); val != 6 {
		t.Errorf("Expected nil at index 4, but got %v", val)
	}
}

func TestLinkedListPrepend(t *testing.T) {
	ll := NewSinglyList()

	ll.Prepend(1)
	if ll.Length != 1 {
		t.Errorf("Expected Length to be 1, got %d", ll.Length)
	}
	if ll.Head.Val != 1 {
		t.Errorf("Expected Head value to be 1, got %d", ll.Head.Val)
	}
	t.Logf("TestLinkedListPrepend passed")
}

func TestLinkedListAppend(t *testing.T) {
	ll := NewSinglyList()

	ll.Append(1)
	if ll.Length != 1 {
		t.Errorf("Expected Length to be 1, got %d", ll.Length)
	}
	if ll.Tail.Val != 1 {
		t.Errorf("Expected Tail value to be 1, got %d", ll.Tail.Val)
	}
}

func TestLinkedListRemove(t *testing.T) {
	ll := NewSinglyList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	out := ll.Remove(2)
	if out != 2 {
		t.Errorf("Expected Remove(2) to return 2, got %d", out)
	}
	if ll.Length != 2 {
		t.Errorf("Expected Length to be 2 after removing, got %d", ll.Length)
	}
}

func TestLinkedListGet(t *testing.T) {
	ll := NewSinglyList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	val := ll.Get(1)
	if val != 2 {
		t.Errorf("Expected Get(1) to return 2, got %d", val)
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	ll := NewSinglyList()

	ll.Append(1)
	ll.Append(3)

	ll.InsertAt(2, 1)
	if ll.Length != 3 {
		t.Errorf("Expected Length to be 3 after InsertAt, got %d", ll.Length)
	}
	if ll.Get(1) != 2 {
		t.Errorf("Expected Get(1) to return 2 after InsertAt, got %d", ll.Get(1))
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	ll := NewSinglyList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	out := ll.RemoveAt(1)
	if out != 2 {
		t.Errorf("Expected RemoveAt(1) to return 2, got %d", out)
	}
	if ll.Length != 2 {
		t.Errorf("Expected Length to be 2 after RemoveAt, got %d", ll.Length)
	}
}
