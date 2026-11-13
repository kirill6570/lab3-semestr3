package queue

import (
	"os"
	"testing"
)

func TestQueueAll(t *testing.T) {
	q := New()
	q.Print()
	if !q.IsEmpty() { 
		t.Error("Expected empty") 
	}

	q.Push("10")
	q.Push("20")
	q.Print()
	if q.Size() != 2 { 
		t.Error("Expected 2") 
	}

	if v, ok := q.Front(); !ok || v != "10" { 
		t.Error("Front") 
	}

	q.Pop()
	q.Pop()
	q.Pop()
	if !q.IsEmpty() { 
		t.Error("Expected empty") 
	}

	if _, ok := q.Front(); ok { 
		t.Error("Front empty") 
	}
}

func TestQueueSerialization(t *testing.T) {
	q := New()
	q.Push("A")
	q.Push("B")

	// Текст
	if err := q.SaveToText("q.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newQ := New()
	if err := newQ.LoadFromText("q.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newQ.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newQ.Size())
	}
	os.Remove("q.json")

	// Бинарь
	if err := q.SaveToBinary("q.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newQb := New()
	if err := newQb.LoadFromBinary("q.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newQb.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newQb.Size())
	}
	os.Remove("q.bin")
}