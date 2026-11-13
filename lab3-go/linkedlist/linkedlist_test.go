package linkedlist

import (
	"os"
	"testing"
)

func TestLinkedListAll(t *testing.T) {
	l := New()
	l.Print()

	l.PushFront("A")
	l.PushBack("B")
	l.PushBack("C")
	l.PushFront("D")
	l.Print()

	if l.Size() != 4 {
		t.Errorf("Expected size 4, got %d", l.Size())
	}

	if l.Find("B") == nil {
		t.Error("Find B failed")
	}
	if l.Find("X") != nil {
		t.Error("Find X should be nil")
	}

	l.RemoveValue("A")
	l.RemoveValue("D")
	l.RemoveValue("C")
	l.RemoveValue("X")
	l.RemoveValue("B")

	if l.Size() != 0 {
		t.Error("Expected empty list")
	}

	l.PopFront()
	l.PopBack()

	l.PushBack("Z")
	l.PopBack()
	l.PushBack("Y")
	l.PopFront()
}

func TestLinkedListSerialization(t *testing.T) {
	l := New()
	l.PushBack("A")
	l.PushBack("B")

	if err := l.SaveToText("l.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newL := New()
	if err := newL.LoadFromText("l.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	os.Remove("l.json")

	if err := l.SaveToBinary("l.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newLb := New()
	if err := newLb.LoadFromBinary("l.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	os.Remove("l.bin")

	newErrText := New()
	_ = newErrText.LoadFromText("non_existent.json")
	newErrBin := New()
	_ = newErrBin.LoadFromBinary("non_existent.bin")
}