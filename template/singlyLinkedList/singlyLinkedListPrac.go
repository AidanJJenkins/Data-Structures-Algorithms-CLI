package singlylinkedlist

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func newNode(value interface{}) *Node {
}

func newSinglyList() *LinkedList {
}

func (ll *LinkedList) Prepend(item interface{}) {
}

func (ll *LinkedList) InsertAt(item interface{}, idx int) {
}

func (ll *LinkedList) Append(item interface{}) {
}

func (ll *LinkedList) Remove(item interface{}) interface{} {
}

func (ll *LinkedList) Get(idx interface{}) interface{} {
}

func (ll *LinkedList) RemoveAt(idx int) interface{} {
}

func (ll *LinkedList) PrintList() {
}
