package doublelinkedlist

import (
	"os"
	"testing"
)

func TestDoubleLinkedListAll(t *testing.T) {
	l := New()
	
	// Печать пустого
	l.PrintForward()
	l.PrintBackward()

	if l.Size() != 0 {
		t.Error("Expected size 0")
	}

	// Добавление
	l.PushBack("A")
	l.PushBack("B")
	l.PushFront("C")

	if l.Size() != 3 {
		t.Error("Expected size 3")
	}

	l.PrintForward()
	l.PrintBackward()

	l.RemoveValue("B") // из середины
	l.RemoveValue("C") // из начала
	l.RemoveValue("A") // из конца
	l.RemoveValue("X") // несуществующий

	// Pop на пустом
	l.PopFront()
	l.PopBack()

	if l.Size() != 0 {
		t.Error("Expected size 0 after removing all")
	}
}

func TestDoubleLinkedListSerialization(t *testing.T) {
	l := New()
	l.PushBack("A")
	l.PushBack("B")

	if err := l.SaveToText("dl.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newL := New()
	if err := newL.LoadFromText("dl.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newL.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newL.Size())
	}
	os.Remove("dl.json")

	if err := l.SaveToBinary("dl.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newLb := New()
	if err := newLb.LoadFromBinary("dl.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newLb.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newLb.Size())
	}
	os.Remove("dl.bin")
}