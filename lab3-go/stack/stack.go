package stack

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	head *Node
	size int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value string) {
	s.head = &Node{Data: value, Next: s.head}
	s.size++
}

func (s *Stack) Pop() (string, bool) {
	if s.head == nil {
		return "", false
	}
	value := s.head.Data
	s.head = s.head.Next
	s.size--
	return value, true
}

func (s *Stack) Peek() (string, bool) {
	if s.head == nil {
		return "", false
	}
	return s.head.Data, true
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("стек пуст")
		return
	}
	fmt.Print("вершина -> ")
	cur := s.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println("<- дно")
}