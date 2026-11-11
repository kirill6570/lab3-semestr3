package array

import (
	"os"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	a := New(10)
	
	if a.Size() != 0 {
		t.Errorf("Ожидался размер 0, получен %d", a.Size())
	}

	a.PushBack("first")
	a.PushBack("second")
	a.PushBack("third")

	if a.Size() != 3 {
		t.Errorf("Ожидался размер 3, получен %d", a.Size())
	}

	if a.Get(0) != "first" || a.Get(1) != "second" {
		t.Errorf("Ошибка получения элемента по индексу")
	}

	a.Set(1, "new")
	if a.Get(1) != "new" {
		t.Errorf("Ошибка обновления элемента")
	}

	a.Set(10, "fail") 
	if a.Get(10) != "" {
		t.Errorf("Ожидалась пустая строка при выходе за границы")
	}

	a.RemoveAt(1)
	if a.Size() != 2 || a.Get(1) != "third" {
		t.Errorf("Ошибка удаления элемента")
	}
	
	a.RemoveAt(10)
	a.Print()
}

func TestSerialization(t *testing.T) {
	da := New(10)
	da.PushBack("hello")
	da.PushBack("world")

	txtFile := "test_array.json"
	err := da.SaveToText(txtFile)
	if err != nil {
		t.Fatalf("Ошибка сохранения в текст: %v", err)
	}

	newDaText := New(10)
	err = newDaText.LoadFromText(txtFile)
	if err != nil {
		t.Fatalf("Ошибка загрузки из текста: %v", err)
	}

	if newDaText.Size() != 2 || newDaText.Get(0) != "hello" || newDaText.Get(1) != "world" {
		t.Errorf("Текстовая десериализация не удалась")
	}
	
	os.Remove(txtFile)

	binFile := "test_array.bin"
	err = da.SaveToBinary(binFile)
	if err != nil {
		t.Fatalf("Ошибка сохранения в бинарь: %v", err)
	}

	newDaBin := New(10)
	err = newDaBin.LoadFromBinary(binFile)
	if err != nil {
		t.Fatalf("Ошибка загрузки из бинаря: %v", err)
	}

	if newDaBin.Size() != 2 || newDaBin.Get(0) != "hello" || newDaBin.Get(1) != "world" {
		t.Errorf("Бинарная десериализация не удалась")
	}

	os.Remove(binFile)
}