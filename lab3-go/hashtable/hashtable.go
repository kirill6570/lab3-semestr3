package hashtable

import "fmt"

const TableSize = 1009

type Node struct {
	Key  string
	Next *Node
}

type HashTable struct {
	buckets [TableSize]*Node
	count   int
}

func New() *HashTable {
	return &HashTable{}
}

func hashString(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h*128 + int(s[i])) % TableSize
	}
	return h
}

func (h *HashTable) Add(key string) {
	idx := hashString(key)
	cur := h.buckets[idx]
	for cur != nil {
		if cur.Key == key {
			return
		}
		cur = cur.Next
	}
	h.buckets[idx] = &Node{Key: key, Next: h.buckets[idx]}
	h.count++
}

func (h *HashTable) Remove(key string) {
	idx := hashString(key)
	var prev *Node
	cur := h.buckets[idx]
	for cur != nil {
		if cur.Key == key {
			if prev != nil {
				prev.Next = cur.Next
			} else {
				h.buckets[idx] = cur.Next
			}
			h.count--
			return
		}
		prev = cur
		cur = cur.Next
	}
}

func (h *HashTable) Contains(key string) bool {
	idx := hashString(key)
	cur := h.buckets[idx]
	for cur != nil {
		if cur.Key == key {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (h *HashTable) Size() int {
	return h.count
}

func (h *HashTable) Print() {
	fmt.Print("{")
	first := true
	for i := 0; i < TableSize; i++ {
		cur := h.buckets[i]
		for cur != nil {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(cur.Key)
			first = false
			cur = cur.Next
		}
	}
	fmt.Println("}")
}