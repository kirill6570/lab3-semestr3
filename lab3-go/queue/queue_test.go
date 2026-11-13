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
	
	q.Push("A")
	q.Push("B")
	q.Push("C")
	
	q.Print()

	if q.Size() != 3 { 
		t.Error("Expected size 3") 
	}
	
	
	if v, ok := q.Front(); !ok || v != "A" { 
		t.Error("Front failed") 
	}
	
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop() // Pop на пустой
	
	if !q.IsEmpty() { 
		t.Error("Expected empty after pops") 
	}
	
	// Front на пустой
	if _, ok := q.Front(); ok {
		t.Error("Front on empty should be false")
	}
}

func TestQueueSerialization(t *testing.T) {
	q := New()
	q.Push("A")
	q.Push("B")

	if err := q.SaveToText("q.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newQ := New()
	if err := newQ.LoadFromText("q.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	os.Remove("q.json")

	if err := q.SaveToBinary("q.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newQb := New()
	if err := newQb.LoadFromBinary("q.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	os.Remove("q.bin")

	newErrText := New()
	_ = newErrText.LoadFromText("non_existent.json")
	newErrBin := New()
	_ = newErrBin.LoadFromBinary("non_existent.bin")
}