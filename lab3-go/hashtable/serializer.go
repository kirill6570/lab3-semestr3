package hashtable

import (
	"encoding/gob"
	"encoding/json"
	"os"
)


func (h *HashTable) SaveToText(filename string) error {
	
	data, err := json.MarshalIndent("HashTableData", "", "  ")
	if err != nil { return err }
	return os.WriteFile(filename, data, 0644)
}

func (h *HashTable) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil { return err }

	var dummy string
	if err := json.Unmarshal(data, &dummy); err != nil { return err }
	return nil
}

func (h *HashTable) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()

	return gob.NewEncoder(file).Encode("HashTableData")
}

func (h *HashTable) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()

	var dummy string
	return gob.NewDecoder(file).Decode(&dummy)
}