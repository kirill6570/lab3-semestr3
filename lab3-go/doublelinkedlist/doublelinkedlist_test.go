package doublelinkedlist

import "testing"

func TestDoubleLinkedListAll(t *testing.T) {
	l := New()
	l.PrintForward()
	l.PrintBackward()

	l.PushFront("a")
	l.PushBack("c")
	l.PushFront("b") // b, a, c
	l.PrintForward()
	l.PrintBackward()
	if l.Size() != 3 {
		t.Error("Expected 3")
	}

	l.PopFront() // a, c
	l.PopBack()  // a
	if l.Size() != 1 {
		t.Error("Expected 1")
	}

	l.RemoveValue("a") // 0
	if l.Size() != 0 {
		t.Error("Expected 0")
	}

	l.PopFront()
	l.PopBack()
	l.RemoveValue("x")
}	