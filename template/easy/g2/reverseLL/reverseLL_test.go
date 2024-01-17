package reversell

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	list := NewSinglyList()
	list.Append(1)
	list.Append(9)
	list.Append(6)
	list.Append(67)

	comp := []any{67, 6, 9, 1}

	list.Reverse()
	linear := printList(list.Head)

	if !reflect.DeepEqual(linear, comp) {
		t.Errorf("lists do not match, reversed: %v, test list: %v", linear, comp)
	}
}

func printList(head *Node) []any {
	var r []any

	curr := head
	for curr != nil {
		r = append(r, curr.Val)
		curr = curr.Next
	}
	return r
}

func TestRec(t *testing.T) {
	list := NewSinglyList()
	list.Append(1)
	list.Append(9)
	list.Append(6)
	list.Append(67)

	comp := []any{67, 6, 9, 1}

	list.Head = Rec(list.Head)
	printed := printList(list.Head)

	if !reflect.DeepEqual(printed, comp) {
		t.Errorf("lists do not match, reversed: %v, test list: %v", printed, comp)
	}

}

func (ll *LinkedList) Append(value int) {
	n := NewNode(value)
	ll.Length++

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
		return
	}

	ll.Tail.Next = n
	ll.Tail = n
}
