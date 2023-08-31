package queue

type val any

type Node struct {
	value val
	next  *Node
}

type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func (q *Queue) Enqueue(item val) {
}

func (q *Queue) Dequeue() any {
}

func (q *Queue) Peek() any {
}
