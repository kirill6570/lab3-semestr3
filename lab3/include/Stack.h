#ifndef STACK_H
#define STACK_H

#include <string>
#include <iostream>

class Stack {
private:
    struct Node {
        std::string data;
        Node* next;
        Node(const std::string& d, Node* n) : data(d), next(n) {}
    };

    Node* head;
    int count;

public:
    Stack() : head(nullptr), count(0) {}
    ~Stack();

    Stack(const Stack&) = delete;
    Stack& operator=(const Stack&) = delete;

    void push(const std::string& value);
    void pop();
    std::string peek() const;
    int size() const;
    bool isEmpty() const;
};

inline Stack::~Stack() {
    while (head != nullptr) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
}

inline void Stack::push(const std::string& value) {
    head = new Node(value, head);
    count++;
}

inline void Stack::pop() {
    if (head == nullptr) return;
    Node* tmp = head;
    head = head->next;
    delete tmp;
    count--;
}

inline std::string Stack::peek() const {
    if (head == nullptr) return "";
    return head->data;
}

inline int Stack::size() const { return count; }

inline bool Stack::isEmpty() const { return head == nullptr; }

#endif