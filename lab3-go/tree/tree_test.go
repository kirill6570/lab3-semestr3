package tree

import (
	"os"
	"testing"
)

func TestCBT(t *testing.T) {
	tr := New()
	if tr.Size() != 0 {
		t.Error("Expected size 0")
	}
	if !tr.IsComplete() {
		t.Error("Empty tree should be complete")
	}

	tr.Insert(10)
	tr.Insert(20)
	tr.Insert(30)
	tr.Insert(40)
	if tr.Size() != 4 {
		t.Error("Expected size 4")
	}
	if !tr.Find(20) || !tr.Find(40) {
		t.Error("Find failed")
	}
	if tr.Find(99) {
		t.Error("Should not find 99")
	}
	if !tr.IsComplete() {
		t.Error("Tree should be complete")
	}

	tr.PrintByLevels()
}

func TestTreeSerialization(t *testing.T) {
	tr := New()
	tr.Insert(10)
	tr.Insert(20)
	tr.Insert(30)


	if err := tr.SaveToText("tree.json"); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newTr := New()
	if err := newTr.LoadFromText("tree.json"); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newTr.Size() != 3 {
		t.Errorf("Expected size 3, got %d", newTr.Size())
	}
	os.Remove("tree.json")

	// Бинарь
	if err := tr.SaveToBinary("tree.bin"); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newTrb := New()
	if err := newTrb.LoadFromBinary("tree.bin"); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newTrb.Size() != 3 {
		t.Errorf("Expected size 3, got %d", newTrb.Size())
	}
	os.Remove("tree.bin")
}