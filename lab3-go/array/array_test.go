package array

import (
	"os"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	a := New(0)
	if a.Size() != 0 {
		t.Error("Expected size 0")
	}

	a.PushBack("first")
	a.PushBack("second")
	a.PushBack("third")

	if a.Size() != 3 {
		t.Errorf("Expected size 3, got %d", a.Size())
	}

	if a.Get(0) != "first" || a.Get(1) != "second" {
		t.Error("Get failed")
	}
	a.Set(1, "new")
	if a.Get(1) != "new" {
		t.Error("Set failed")
	}

	// Выход за границы
	a.Set(10, "fail")
	a.Get(10)
	a.Get(-1)
	a.RemoveAt(10)
	a.RemoveAt(-1)

	a.RemoveAt(1)
	if a.Size() != 2 {
		t.Error("RemoveAt failed")
	}

	a.PopBack()
	a.PopBack()
	a.PopBack() // на пустом
	if a.Size() != 0 {
		t.Error("Expected empty")
	}

	a.Print()

	// Проверяем Print на заполненном
	a.PushBack("A")
	a.Print()
}

func TestNewWithItem(t *testing.T) {
	a := NewWithItem(3, "x")
	if a.Size() != 3 {
		t.Errorf("Expected size 3, got %d", a.Size())
	}
	
	b := NewWithItem(-5, "y")
	if b.Size() != 1 {
		t.Errorf("Expected size 1, got %d", b.Size())
	}
}

func TestSerialization(t *testing.T) {
	da := New(10)
	da.PushBack("hello")
	da.PushBack("world")

	// Текст
	txtFile := "test_array.json"
	if err := da.SaveToText(txtFile); err != nil {
		t.Fatalf("Text save error: %v", err)
	}
	newDaText := New(10)
	if err := newDaText.LoadFromText(txtFile); err != nil {
		t.Fatalf("Text load error: %v", err)
	}
	if newDaText.Size() != 2 {
		t.Error("Text load failed")
	}
	os.Remove(txtFile)

	// Бинарь
	binFile := "test_array.bin"
	if err := da.SaveToBinary(binFile); err != nil {
		t.Fatalf("Binary save error: %v", err)
	}
	newDaBin := New(10)
	if err := newDaBin.LoadFromBinary(binFile); err != nil {
		t.Fatalf("Binary load error: %v", err)
	}
	if newDaBin.Size() != 2 {
		t.Error("Binary load failed")
	}
	os.Remove(binFile)
	newErrText := New(10)
	_ = newErrText.LoadFromText("non_existent.json")
	
	newErrBin := New(10)
	_ = newErrBin.LoadFromBinary("non_existent.bin")
}