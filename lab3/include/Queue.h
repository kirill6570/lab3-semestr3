#ifndef QUEUE_H
#define QUEUE_H

#include <string>
#include <iostream>

class Queue {
private:
    struct Node {
        std::string data;
        Node* next;
        Node(const std::string& d, Node* n) : data(d), next(n) {}
    };

    Node* head;
    Node* tail;
    int count;

public:
    Queue() : head(nullptr), tail(nullptr), count(0) {}
    ~Queue();

    Queue(const Queue&) = delete;
    Queue& operator=(const Queue&) = delete;

    void push(const std::string& value);
    void pop();
    std::string front() const;
    int size() const;
    bool isEmpty() const;
};

inline Queue::~Queue() {
    while (head != nullptr) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
}

inline void Queue::push(const std::string& value) {
    Node* node = new Node(value, nullptr);
    if (tail == nullptr) {
        head = node;
        tail = node;
    } else {
        tail->next = node;
        tail = node;
    }
    count++;
}

inline void Queue::pop() {
    if (head == nullptr) return;
    Node* tmp = head;
    head = head->next;
    if (head == nullptr) tail = nullptr;
    delete tmp;
    count--;
}

inline std::string Queue::front() const {
    if (head == nullptr) return "";
    return head->data;
}

inline int Queue::size() const { return count; }

inline bool Queue::isEmpty() const { return head == nullptr; }

#endif