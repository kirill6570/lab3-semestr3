package stack

import "testing"

func TestStackAll(t *testing.T) {
	s := New()
	s.Print()
	if !s.IsEmpty() { t.Error("Expected empty") }

	s.Push("10")
	s.Push("20")
	s.Print()
	if s.Size() != 2 { t.Error("Expected 2") }

	if v, ok := s.Peek(); !ok || v != "20" { t.Error("Peek") }

	s.Pop()
	s.Pop()
	s.Pop()
	if !s.IsEmpty() { t.Error("Expected empty") }

	if _, ok := s.Peek(); ok { t.Error("Peek empty") }
}