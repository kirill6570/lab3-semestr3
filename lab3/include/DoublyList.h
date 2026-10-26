#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <string>
#include <iostream>

class DoublyList {
private:
    struct Node {
        std::string data;
        Node* prev;
        Node* next;
        Node(const std::string& d, Node* p, Node* n) : data(d), prev(p), next(n) {}
    };

    Node* head;
    Node* tail;
    int size;

public:
    DoublyList() : head(nullptr), tail(nullptr), size(0) {}
    ~DoublyList();

    DoublyList(const DoublyList&) = delete;
    DoublyList& operator=(const DoublyList&) = delete;

    void pushFront(const std::string& value);
    void pushBack(const std::string& value);
    void popFront();
    void popBack();
    bool find(const std::string& value) const;
    void remove(const std::string& value);
    int getSize() const;
    void printForward() const;
    void printBackward() const;
};

inline DoublyList::~DoublyList() {
    while (head != nullptr) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
}

inline void DoublyList::pushFront(const std::string& value) {
    Node* node = new Node(value, nullptr, head);
    if (head != nullptr) head->prev = node;
    else tail = node;
    head = node;
    size++;
}

inline void DoublyList::pushBack(const std::string& value) {
    Node* node = new Node(value, tail, nullptr);
    if (tail != nullptr) tail->next = node;
    else head = node;
    tail = node;
    size++;
}

inline void DoublyList::popFront() {
    if (head == nullptr) return;
    Node* tmp = head;
    head = head->next;
    if (head != nullptr) head->prev = nullptr;
    else tail = nullptr;
    delete tmp;
    size--;
}

inline void DoublyList::popBack() {
    if (tail == nullptr) return;
    Node* tmp = tail;
    tail = tail->prev;
    if (tail != nullptr) tail->next = nullptr;
    else head = nullptr;
    delete tmp;
    size--;
}

inline bool DoublyList::find(const std::string& value) const {
    Node* cur = head;
    while (cur != nullptr) {
        if (cur->data == value) return true;
        cur = cur->next;
    }
    return false;
}

inline void DoublyList::remove(const std::string& value) {
    Node* cur = head;
    while (cur != nullptr) {
        if (cur->data == value) {
            if (cur->prev != nullptr) cur->prev->next = cur->next;
            else head = cur->next;
            if (cur->next != nullptr) cur->next->prev = cur->prev;
            else tail = cur->prev;
            delete cur;
            size--;
            return;
        }
        cur = cur->next;
    }
}

inline int DoublyList::getSize() const { return size; }

inline void DoublyList::printForward() const {
    Node* cur = head;
    while (cur != nullptr) {
        std::cout << cur->data << " ";
        cur = cur->next;
    }
    std::cout << "\n";
}

inline void DoublyList::printBackward() const {
    Node* cur = tail;
    while (cur != nullptr) {
        std::cout << cur->data << " ";
        cur = cur->prev;
    }
    std::cout << "\n";
}

#endif