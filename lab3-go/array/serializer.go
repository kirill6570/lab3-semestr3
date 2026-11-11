package array

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

// --- ТЕКСТОВАЯ (JSON) ---
func (a *DynamicArray) SaveToText(filename string) error {
	// Используем анонимную структуру, чтобы обойти приватность поля "data"
	exported := struct {
		Data []string `json:"data"`
	}{Data: a.data}

	data, err := json.MarshalIndent(exported, "", "  ")
	if err != nil { return err }
	return os.WriteFile(filename, data, 0644)
}

func (a *DynamicArray) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil { return err }

	exported := struct {
		Data []string `json:"data"`
	}{}
	
	err = json.Unmarshal(data, &exported)
	if err != nil { return err }
	
	a.data = exported.Data
	return nil
}

// --- БИНАРНАЯ (GOB) ---
func (a *DynamicArray) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()

	// Для gob тоже используем временную структуру
	exported := struct{ Data []string }{Data: a.data}
	return gob.NewEncoder(file).Encode(exported)
}

func (a *DynamicArray) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()

	exported := struct{ Data []string }{}
	err = gob.NewDecoder(file).Decode(&exported)
	if err != nil { return err }

	a.data = exported.Data
	return nil
}