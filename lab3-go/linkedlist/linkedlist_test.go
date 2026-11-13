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

	// 3. Поиск
	if l.Find("B") == nil {
		t.Error("Find B failed")
	}
	if l.Find("X") != nil {
		t.Error("Find X should be nil")
	}


	l.RemoveValue("A") // Из середины
	l.RemoveValue("D") // Из начала
	l.RemoveValue("C")
	l.RemoveValue("X") // Несуществующий
	l.RemoveValue("B") // Последний элемент
	
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

	// Текст
	if err := l.SaveToText("l.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newL := New()
	if err := newL.LoadFromText("l.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newL.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newL.Size())
	}
	os.Remove("l.json")

	// Бинарь
	if err := l.SaveToBinary("l.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newLb := New()
	if err := newLb.LoadFromBinary("l.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newLb.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newLb.Size())
	}
	os.Remove("l.bin")
}