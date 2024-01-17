package queue

type QueueNode struct {
	Val  int
	Next *QueueNode
}

type Queue struct {
	Length     int
	Head, Tail *QueueNode
}

func NewQueue() *Queue {
}

func (q *Queue) Enqueue(value int) {

}

func (q *Queue) Dequeue() any {

}

func (q *Queue) Peek() any {
}
