package array

import "testing"

func TestDynamicArray(t *testing.T) {
	a := New(4)
	if a.Size() != 0 {
		t.Error("Expected size 0")
	}

	a.PushBack("10")
	a.PushBack("20")
	a.PushBack("30")
	if a.Size() != 3 {
		t.Error("Expected size 3")
	}
	if a.Get(1) != "20" {
		t.Error("Expected 20 at index 1")
	}

	a.Set(1, "99")
	if a.Get(1) != "99" {
		t.Error("Expected 99")
	}

	a.RemoveAt(1)
	if a.Size() != 2 || a.Get(1) != "30" {
		t.Error("Remove failed")
	}

	a.PopBack()
	if a.Size() != 1 {
		t.Error("PopBack failed")
	}

	a.Print()
}

func TestNewWithItem(t *testing.T) {
	a := NewWithItem(3, "x")
	if a.Size() != 3 {
		t.Error("Expected size 3")
	}
	if a.Get(0) != "x" || a.Get(2) != "x" {
		t.Error("Expected all x")
	}
}

func TestOutOfBounds(t *testing.T) {
	a := New(2)
	if a.Get(-1) != "" || a.Get(100) != "" {
		t.Error("Out of bounds should return empty")
	}
	a.Set(-1, "x")
	a.RemoveAt(-1)
	if a.Size() != 0 {
		t.Error("Invalid ops should not change size")
	}
}