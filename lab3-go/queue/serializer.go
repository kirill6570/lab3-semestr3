package queue

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

func (q *Queue) SaveToText(filename string) error {
	// Используем методы твоей очереди, чтобы достать данные
	var values []string
	tempQ := New() // создаем временную очередь
	
	// Выгребаем все элементы из копии (чтобы не испортить оригинал)
	for !q.IsEmpty() {
		v, _ := q.Pop() 
		values = append(values, v)
		tempQ.Push(v) // сохраняем во временную
	}
	
	// Возвращаем элементы обратно в оригинал
	for !tempQ.IsEmpty() {
		v, _ := tempQ.Pop()
		q.Push(v)
	}

	data, err := json.MarshalIndent(values, "", "  ")
	if err != nil { return err }
	return os.WriteFile(filename, data, 0644)
}

func (q *Queue) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil { return err }

	var values []string
	if err := json.Unmarshal(data, &values); err != nil { return err }

	// Очищаем текущую очередь (выгребаем всё)
	for !q.IsEmpty() {
		q.Pop()
	}
	
	// Загружаем новые
	for _, v := range values {
		q.Push(v)
	}
	return nil
}

func (q *Queue) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()

	var values []string
	tempQ := New()
	for !q.IsEmpty() {
		v, _ := q.Pop()
		values = append(values, v)
		tempQ.Push(v)
	}
	for !tempQ.IsEmpty() {
		v, _ := tempQ.Pop()
		q.Push(v)
	}

	return gob.NewEncoder(file).Encode(values)
}

func (q *Queue) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()

	var values []string
	if err := gob.NewDecoder(file).Decode(&values); err != nil { return err }

	for !q.IsEmpty() {
		q.Pop()
	}
	for _, v := range values {
		q.Push(v)
	}
	return nil
}