package linkedlist

import "testing"

func TestLinkedListAll(t *testing.T) {
	l := New()
	l.Print()

	if l.Size() != 0 {
		t.Error("Expected size 0")
	}

	l.PushFront("a")
	l.PushBack("b")
	l.PushBack("c")
	l.Print()
	if l.Size() != 3 {
		t.Error("Expected size 3")
	}

	if l.Find("a") == nil {
		t.Error("Should find a")
	}
	if l.Find("z") != nil {
		t.Error("Should not find z")
	}

	// Удаление головы
	l.RemoveValue("a")
	if l.Size() != 2 {
		t.Error("Remove head failed")
	}

	// Удаление хвоста
	l.RemoveValue("c")
	if l.Size() != 1 {
		t.Error("Remove tail failed")
	}

	// Удаление середины
	l.PushBack("d")
	l.PushBack("e")
	l.RemoveValue("d")
	if l.Size() != 2 {
		t.Error("Remove middle failed")
	}

	l.PopFront()
	if l.Size() != 1 {
		t.Error("PopFront failed")
	}

	// PopBack (единственный элемент)
	l.PopBack()
	if l.Size() != 0 {
		t.Error("PopBack single failed")
	}

	// Пустые операции
	l.PopFront()
	l.PopBack()
	l.RemoveValue("x")

	// PopFront единственного
	l.PushFront("only")
	l.PopFront()
	if l.Size() != 0 {
		t.Error("PopFront single failed")
	}
}
func TestPopBackSingle(t *testing.T) {
	l := New()
	l.PushBack("only")
	l.PopBack()
	if l.Size() != 0 {
		t.Error("Expected 0 after single PopBack")
	}
	if l.Find("only") != nil {
		t.Error("Should be removed")
	}
}