package hashtable

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

func (h *HashTable) toSlice() []string {
	var res []string
	for i := 0; i < TableSize; i++ {
		cur := h.buckets[i]
		for cur != nil {
			res = append(res, cur.Key)
			cur = cur.Next
		}
	}
	return res
}

func (h *HashTable) SaveToText(filename string) error {
	data, err := json.MarshalIndent(h.toSlice(), "", "  ")
	if err != nil { return err }
	return os.WriteFile(filename, data, 0644)
}

func (h *HashTable) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil { return err }

	var keys []string
	if err := json.Unmarshal(data, &keys); err != nil { return err }

	h.buckets = [TableSize]*Node{}
	h.count = 0
	for _, k := range keys {
		h.Add(k)
	}
	return nil
}

func (h *HashTable) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()
	return gob.NewEncoder(file).Encode(h.toSlice())
}

func (h *HashTable) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()

	var keys []string
	if err := gob.NewDecoder(file).Decode(&keys); err != nil { return err }

	h.buckets = [TableSize]*Node{}
	h.count = 0
	for _, k := range keys {
		h.Add(k)
	}
	return nil
}