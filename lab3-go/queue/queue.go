package queue

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(value string) {
	node := &Node{Data: value}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) Pop() (string, bool) {
	if q.head == nil {
		return "", false
	}
	value := q.head.Data
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, true
}

func (q *Queue) Front() (string, bool) {
	if q.head == nil {
		return "", false
	}
	return q.head.Data, true
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Print() {
	if q.IsEmpty() {
		fmt.Println("очередь пуста")
		return
	}
	fmt.Print("начало -> ")
	cur := q.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println("<- конец")
}