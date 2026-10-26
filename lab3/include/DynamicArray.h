#ifndef DYNAMICARRAY_H
#define DYNAMICARRAY_H

#include <string>
#include <iostream>

class DynamicArray {
private:
    std::string* data;
    int size;
    int capacity;

    void ensureCapacity(int needed);

public:
    DynamicArray(int initialCapacity = 4)
        : data(new std::string[initialCapacity < 1 ? 1 : initialCapacity]),
          size(0),
          capacity(initialCapacity < 1 ? 1 : initialCapacity) {}

    ~DynamicArray() { delete[] data; }

    // Запрет копирования
    DynamicArray(const DynamicArray&) = delete;
    DynamicArray& operator=(const DynamicArray&) = delete;

    // Перемещение
    DynamicArray(DynamicArray&& other) noexcept
        : data(other.data), size(other.size), capacity(other.capacity) {
        other.data = nullptr;
        other.size = 0;
        other.capacity = 0;
    }

    DynamicArray& operator=(DynamicArray&& other) noexcept {
        if (this != &other) {
            delete[] data;
            data = other.data;
            size = other.size;
            capacity = other.capacity;
            other.data = nullptr;
            other.size = 0;
            other.capacity = 0;
        }
        return *this;
    }

    void pushBack(const std::string& value) {
        ensureCapacity(size + 1);
        data[size++] = value;
    }

    void popBack() { if (size > 0) size--; }

    std::string get(int index) const {
        if (index < 0 || index >= size) return "";
        return data[index];
    }

    void set(int index, const std::string& value) {
        if (index < 0 || index >= size) return;
        data[index] = value;
    }

    void removeAt(int index) {
        if (index < 0 || index >= size) return;
        for (int i = index; i < size - 1; ++i) data[i] = data[i + 1];
        size--;
    }

    int getSize() const { return size; }

    void print() const {
        std::cout << "[";
        for (int i = 0; i < size; ++i) {
            std::cout << data[i];
            if (i < size - 1) std::cout << ", ";
        }
        std::cout << "]\n";
    }
};

inline void DynamicArray::ensureCapacity(int needed) {
    if (needed <= capacity) return;
    int newCap = capacity;
    while (newCap < needed) newCap *= 2;
    std::string* newData = new std::string[newCap];
    for (int i = 0; i < size; ++i) newData[i] = data[i];
    delete[] data;
    data = newData;
    capacity = newCap;
}

#endif