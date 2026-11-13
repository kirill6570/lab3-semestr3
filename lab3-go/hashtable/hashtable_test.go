package hashtable

import (
	"os"
	"testing"
)

func TestHashTableAll(t *testing.T) {
	h := New()
	if h.Size() != 0 {
		t.Error("Expected size 0")
	}

	h.Add("apple")
	h.Add("banana")
	h.Add("cherry")
	h.Add("apple")

	if h.Size() != 3 {
		t.Errorf("Expected size 3, got %d", h.Size())
	}

	if !h.Contains("apple") {
		t.Error("Contains failed for apple")
	}
	if h.Contains("grape") {
		t.Error("Contains failed for grape")
	}

	h.Remove("banana")
	if h.Size() != 2 {
		t.Errorf("Expected size 2 after remove, got %d", h.Size())
	}
	h.Print()
}

func TestHashTableSerialization(t *testing.T) {
	h := New()
	h.Add("apple")
	h.Add("banana")

	if err := h.SaveToText("ht.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newH := New()
	if err := newH.LoadFromText("ht.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	os.Remove("ht.json")

	if err := h.SaveToBinary("ht.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newHb := New()
	if err := newHb.LoadFromBinary("ht.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	os.Remove("ht.bin")

	newErrText := New()
	_ = newErrText.LoadFromText("non_existent.json")
	newErrBin := New()
	_ = newErrBin.LoadFromBinary("non_existent.bin")
}