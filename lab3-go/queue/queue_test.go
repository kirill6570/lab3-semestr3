package queue

import "testing"

func TestQueueAll(t *testing.T) {
	q := New()
	q.Print()
	if !q.IsEmpty() { t.Error("Expected empty") }

	q.Push("10")
	q.Push("20")
	q.Print()
	if q.Size() != 2 { t.Error("Expected 2") }

	if v, ok := q.Front(); !ok || v != "10" { t.Error("Front") }

	q.Pop()
	q.Pop()
	q.Pop()
	if !q.IsEmpty() { t.Error("Expected empty") }

	if _, ok := q.Front(); ok { t.Error("Front empty") }
}