package stack

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

// Собираем стек в обычный слайс для сериализации
func (s *Stack) toSlice() []string {
	var res []string
	cur := s.head
	for cur != nil {
		res = append(res, cur.Data)
		cur = cur.Next
	}
	return res
}

// --- ТЕКСТОВАЯ СЕРИАЛИЗАЦИЯ (JSON) ---

func (s *Stack) SaveToText(filename string) error {
	data, err := json.MarshalIndent(s.toSlice(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (s *Stack) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var values []string
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	s.head = nil
	s.size = 0
	for _, v := range values {
		s.Push(v)
	}
	return nil
}

// --- БИНАРНАЯ СЕРИАЛИЗАЦИЯ (GOB) ---

func (s *Stack) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(s.toSlice())
}

func (s *Stack) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var values []string
	if err := gob.NewDecoder(file).Decode(&values); err != nil {
		return err
	}

	s.head = nil
	s.size = 0
	for _, v := range values {
		s.Push(v)
	}
	return nil
}