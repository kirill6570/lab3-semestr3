package stack

import (
	"os"
	"testing"
)

func TestStackAll(t *testing.T) {
	s := New()
	s.Print()
	if !s.IsEmpty() { 
		t.Error("Expected empty") 
	}

	s.Push("10")
	s.Push("20")
	s.Print()
	if s.Size() != 2 { 
		t.Error("Expected 2") 
	}

	if v, ok := s.Peek(); !ok || v != "20" { 
		t.Error("Peek") 
	}

	s.Pop()
	s.Pop()
	s.Pop()
	if !s.IsEmpty() { 
		t.Error("Expected empty") 
	}

	if _, ok := s.Peek(); ok { 
		t.Error("Peek empty") 
	}
}

func TestStackSerialization(t *testing.T) {
	s := New()
	s.Push("A")
	s.Push("B")

	// Текст
	if err := s.SaveToText("s.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newS := New()
	if err := newS.LoadFromText("s.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newS.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newS.Size())
	}
	os.Remove("s.json")

	// Бинарь
	if err := s.SaveToBinary("s.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newSb := New()
	if err := newSb.LoadFromBinary("s.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newSb.Size() != 2 {
		t.Errorf("Expected size 2, got %d", newSb.Size())
	}
	os.Remove("s.bin")
}