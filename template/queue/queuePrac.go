package queue

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func (q *Queue) Enqueue(item any) {
}

func (q *Queue) Dequeue() any {
}

func (q *Queue) Peek() any {
}
