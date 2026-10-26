package hashtable

import "testing"

func TestHashTable(t *testing.T) {
	h := New()
	if h.Size() != 0 {
		t.Error("Expected size 0")
	}
	if h.Contains("apple") {
		t.Error("Should not contain apple")
	}

	h.Add("apple")
	h.Add("banana")
	h.Add("cherry")
	if h.Size() != 3 {
		t.Error("Expected size 3")
	}
	if !h.Contains("apple") || !h.Contains("banana") {
		t.Error("Contains failed")
	}

	h.Add("apple")
	if h.Size() != 3 {
		t.Error("Duplicate added")
	}

	h.Remove("banana")
	if h.Size() != 2 {
		t.Error("Expected size 2")
	}
	if h.Contains("banana") {
		t.Error("Banana should be removed")
	}

	h.Remove("nonexistent")
	if h.Size() != 2 {
		t.Error("Remove nonexistent should not change size")
	}

	h.Print()
}