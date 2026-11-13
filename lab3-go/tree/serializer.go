package tree

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

func (t *CompleteBinaryTree) toSlice() []int {
	var res []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Data)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(t.root)
	return res
}

func (t *CompleteBinaryTree) SaveToText(filename string) error {
	data, err := json.MarshalIndent(t.toSlice(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *CompleteBinaryTree) LoadFromText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var values []int
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	t.root = nil
	t.size = 0
	for _, v := range values {
		t.Insert(v)
	}
	return nil
}

func (t *CompleteBinaryTree) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(t.toSlice())
}

func (t *CompleteBinaryTree) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var values []int
	if err := gob.NewDecoder(file).Decode(&values); err != nil {
		return err
	}

	t.root = nil
	t.size = 0
	for _, v := range values {
		t.Insert(v)
	}
	return nil
}