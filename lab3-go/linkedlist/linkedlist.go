package linkedlist

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) PushFront(value string) {
	node := &Node{Data: value, Next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

func (l *LinkedList) PushBack(value string) {
	node := &Node{Data: value}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.size++
}

func (l *LinkedList) PopFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
}

func (l *LinkedList) PopBack() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}
	cur := l.head
	for cur.Next != l.tail {
		cur = cur.Next
	}
	cur.Next = nil
	l.tail = cur
	l.size--
}

func (l *LinkedList) Find(value string) *Node {
	cur := l.head
	for cur != nil {
		if cur.Data == value {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (l *LinkedList) RemoveValue(value string) {
	if l.head == nil {
		return
	}
	if l.head.Data == value {
		l.PopFront()
		return
	}
	prev := l.head
	cur := l.head.Next
	for cur != nil {
		if cur.Data == value {
			prev.Next = cur.Next
			if cur == l.tail {
				l.tail = prev
			}
			l.size--
			return
		}
		prev = cur
		cur = cur.Next
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("список пуст")
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Print(cur.Data)
		if cur.Next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.Next
	}
	fmt.Println(" -> NULL")
}