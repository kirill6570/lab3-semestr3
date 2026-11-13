package doublelinkedlist

import (
	"os"
	"testing"
)

func TestDoubleLinkedListAll(t *testing.T) {
	l := New()

	l.PrintForward()
	l.PrintBackward()

	if l.Size() != 0 {
		t.Error("Expected size 0")
	}

	l.PushBack("A")
	l.PushBack("B")
	l.PushFront("C")

	if l.Size() != 3 {
		t.Error("Expected size 3")
	}

	l.PrintForward()
	l.PrintBackward()

	l.RemoveValue("X")
	l.RemoveValue("C")
	l.RemoveValue("A")
	l.RemoveValue("B")

	l.PopFront()
	l.PopBack()

	l.PushBack("E")
	l.PushFront("F")

	l.PopFront()
	l.PopFront()

	l.PushBack("G")
	l.PopBack()
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

	newErrText := New()
	_ = newErrText.LoadFromText("non_existent.json")
	newErrBin := New()
	_ = newErrBin.LoadFromBinary("non_existent.bin")
}