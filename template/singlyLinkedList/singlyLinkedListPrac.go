package singlylinkedlist

type value interface{}

type Node struct {
	value value
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func newNode(value value) *Node {
}

func newSinglyList() *LinkedList {
}

func (ll *LinkedList) Prepend(item value) {
}

func (ll *LinkedList) InsertAt(item value, idx int) {
}

func (ll *LinkedList) Append(item value) {
}

func (ll *LinkedList) Remove(item value) interface{} {
}

func (ll *LinkedList) Get(idx value) interface{} {
}

func (ll *LinkedList) RemoveAt(idx int) value {
}
