#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#include <string>
#include <iostream>

class SinglyList {
private:
    struct Node {
        std::string data;
        Node* next;
        Node(const std::string& d, Node* n) : data(d), next(n) {}
    };

    Node* head;
    Node* tail;
    int size;

public:
    SinglyList() : head(nullptr), tail(nullptr), size(0) {}
    ~SinglyList();

    // Запрет копирования
    SinglyList(const SinglyList&) = delete;
    SinglyList& operator=(const SinglyList&) = delete;

    // Перемещение
    SinglyList(SinglyList&& other) noexcept
        : head(other.head), tail(other.tail), size(other.size) {
        other.head = nullptr;
        other.tail = nullptr;
        other.size = 0;
    }

    SinglyList& operator=(SinglyList&& other) noexcept {
        if (this != &other) {
            while (head != nullptr) {
                Node* tmp = head;
                head = head->next;
                delete tmp;
            }
            head = other.head;
            tail = other.tail;
            size = other.size;
            other.head = nullptr;
            other.tail = nullptr;
            other.size = 0;
        }
        return *this;
    }

    void pushFront(const std::string& value);
    void pushBack(const std::string& value);
    void popFront();
    void popBack();
    bool find(const std::string& value) const;
    void remove(const std::string& value);
    int getSize() const;
    std::string get(int index) const;
    void print() const;
};

inline SinglyList::~SinglyList() {
    while (head != nullptr) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
}

inline void SinglyList::pushFront(const std::string& value) {
    Node* node = new Node(value, head);
    head = node;
    if (tail == nullptr) tail = node;
    size++;
}

inline void SinglyList::pushBack(const std::string& value) {
    Node* node = new Node(value, nullptr);
    if (tail == nullptr) {
        head = node;
        tail = node;
    } else {
        tail->next = node;
        tail = node;
    }
    size++;
}

inline void SinglyList::popFront() {
    if (head == nullptr) return;
    Node* tmp = head;
    head = head->next;
    if (head == nullptr) tail = nullptr;
    delete tmp;
    size--;
}

inline void SinglyList::popBack() {
    if (head == nullptr) return;
    if (head == tail) {
        delete head;
        head = nullptr;
        tail = nullptr;
        size = 0;
        return;
    }
    Node* cur = head;
    while (cur->next != tail) cur = cur->next;
    delete tail;
    cur->next = nullptr;
    tail = cur;
    size--;
}

inline bool SinglyList::find(const std::string& value) const {
    Node* cur = head;
    while (cur != nullptr) {
        if (cur->data == value) return true;
        cur = cur->next;
    }
    return false;
}

inline void SinglyList::remove(const std::string& value) {
    if (head == nullptr) return;
    if (head->data == value) {
        popFront();
        return;
    }
    Node* prev = head;
    Node* cur = head->next;
    while (cur != nullptr) {
        if (cur->data == value) {
            prev->next = cur->next;
            if (cur == tail) tail = prev;
            delete cur;
            size--;
            return;
        }
        prev = cur;
        cur = cur->next;
    }
}

inline int SinglyList::getSize() const { return size; }

inline std::string SinglyList::get(int index) const {
    if (index < 0 || index >= size) return "";
    Node* cur = head;
    for (int i = 0; i < index && cur; ++i) cur = cur->next;
    return cur ? cur->data : "";
}

inline void SinglyList::print() const {
    if (head == nullptr) {
        std::cout << "список пуст\n";
        return;
    }
    Node* cur = head;
    while (cur != nullptr) {
        std::cout << cur->data;
        if (cur->next != nullptr) std::cout << " -> ";
        cur = cur->next;
    }
    std::cout << " -> NULL\n";
}

#endif