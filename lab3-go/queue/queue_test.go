package queue

import "testing"

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
		t.Error("Expected size 2") 
	}

	if v, ok := q.Front(); !ok || v != "10" { 
		t.Error("Front should be 10") 
	}

	if v, ok := q.Pop(); !ok || v != "10" {
		t.Error("Pop should return 10")
	}
	if v, ok := q.Pop(); !ok || v != "20" {
		t.Error("Pop should return 20")
	}
	
	q.Pop() 
	
	if !q.IsEmpty() { 
		t.Error("Expected empty after pops") 
	}

	if _, ok := q.Front(); ok { 
		t.Error("Front on empty queue should return false") 
	}
}