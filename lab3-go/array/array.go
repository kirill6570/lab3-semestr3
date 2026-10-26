package array

import "fmt"

type DynamicArray struct {
	data []string
}

func New(initialCapacity int) *DynamicArray {
	if initialCapacity < 1 {
		initialCapacity = 1
	}
	return &DynamicArray{data: make([]string, 0, initialCapacity)}
}

func NewWithItem(initialCapacity int, item string) *DynamicArray {
	if initialCapacity < 1 {
		initialCapacity = 1
	}
	data := make([]string, initialCapacity)
	for i := range data {
		data[i] = item
	}
	return &DynamicArray{data: data}
}

func (a *DynamicArray) PushBack(value string) {
	a.data = append(a.data, value)
}

func (a *DynamicArray) PopBack() {
	if len(a.data) > 0 {
		a.data = a.data[:len(a.data)-1]
	}
}

func (a *DynamicArray) Get(index int) string {
	if index < 0 || index >= len(a.data) {
		return ""
	}
	return a.data[index]
}

func (a *DynamicArray) Set(index int, value string) {
	if index < 0 || index >= len(a.data) {
		return
	}
	a.data[index] = value
}

func (a *DynamicArray) RemoveAt(index int) {
	if index < 0 || index >= len(a.data) {
		return
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
}

func (a *DynamicArray) Size() int {
	return len(a.data)
}

func (a *DynamicArray) Print() {
	fmt.Print("[")
	for i, v := range a.data {
		fmt.Print(v)
		if i < len(a.data)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}