package tree

import "testing"

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