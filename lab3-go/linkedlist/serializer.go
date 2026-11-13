package linkedlist

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

func (l *LinkedList) toSlice() []string {
	var res []string
	cur := l.head 
	for cur != nil {
		res = append(res, cur.Data) 
		cur = cur.Next
	}
	return res
}

func (l *LinkedList) SaveToText(filename string) error {
	data, err := json.MarshalIndent(l.toSlice(), "", "  ")
	if err != nil { return err }
	return os.WriteFile(filename, data, 0644)
}

func (l *LinkedList) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil { return err }

	var values []string
	if err := json.Unmarshal(data, &values); err != nil { return err }

	l.head = nil 
	l.size = 0  
	for _, v := range values {
		l.PushBack(v)
	}
	return nil
}

func (l *LinkedList) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()
	return gob.NewEncoder(file).Encode(l.toSlice())
}

func (l *LinkedList) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()

	var values []string
	if err := gob.NewDecoder(file).Decode(&values); err != nil { return err }

	l.head = nil
	l.size = 0
	for _, v := range values {
		l.PushBack(v)
	}
	return nil
}