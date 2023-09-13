package reversell

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	list := newSinglyList()
	list.Append(1)
	list.Append(9)
	list.Append(6)
	list.Append(67)

	comp := []any{67, 6, 9, 1}

	list.Reverse()
	linear := printList(list.head)

	if !reflect.DeepEqual(linear, comp) {
		t.Errorf("lists do not match, reversed: %v, test list: %v", linear, comp)
	}
}

func printList(head *Node) []any {
	var r []any

	curr := head
	for curr != nil {
		r = append(r, curr.value)
		curr = curr.next
	}
	return r
}

func TestRec(t *testing.T) {
	list := newSinglyList()
	list.Append(1)
	list.Append(9)
	list.Append(6)
	list.Append(67)

	comp := []any{67, 6, 9, 1}

	list.head = Rec(list.head)
	printed := printList(list.head)

	if !reflect.DeepEqual(printed, comp) {
		t.Errorf("lists do not match, reversed: %v, test list: %v", printed, comp)
	}

}

func (ll *LinkedList) Append(item value) {
	n := newNode(item)
	ll.length++

	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}

	ll.tail.next = n
	ll.tail = n
}
