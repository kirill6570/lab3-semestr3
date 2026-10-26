#ifndef HASHTABLE_H
#define HASHTABLE_H

#include <string>
#include <iostream>

class HashTable {
private:
    struct Node {
        std::string key;
        Node* next;
        Node(const std::string& k, Node* n) : key(k), next(n) {}
    };

    static const int TABLE_SIZE = 1009;
    Node* buckets[TABLE_SIZE];
    int count;

    int hashString(const std::string& s) const;

public:
    HashTable();
    ~HashTable();

    HashTable(const HashTable&) = delete;
    HashTable& operator=(const HashTable&) = delete;

    void add(const std::string& key);
    void remove(const std::string& key);
    bool contains(const std::string& key) const;
    int size() const;
};

inline HashTable::HashTable() : count(0) {
    for (int i = 0; i < TABLE_SIZE; ++i) buckets[i] = nullptr;
}

inline HashTable::~HashTable() {
    for (int i = 0; i < TABLE_SIZE; ++i) {
        Node* cur = buckets[i];
        while (cur != nullptr) {
            Node* next = cur->next;
            delete cur;
            cur = next;
        }
    }
}

inline int HashTable::hashString(const std::string& s) const {
    int h = 0;
    for (char c : s) {
        h = (h * 128 + (unsigned char)c) % TABLE_SIZE;
    }
    return h;
}

inline void HashTable::add(const std::string& key) {
    int h = hashString(key);
    Node* cur = buckets[h];
    while (cur != nullptr) {
        if (cur->key == key) return;
        cur = cur->next;
    }
    buckets[h] = new Node(key, buckets[h]);
    count++;
}

inline void HashTable::remove(const std::string& key) {
    int h = hashString(key);
    Node* cur = buckets[h];
    Node* prev = nullptr;
    while (cur != nullptr) {
        if (cur->key == key) {
            if (prev != nullptr) prev->next = cur->next;
            else buckets[h] = cur->next;
            delete cur;
            count--;
            return;
        }
        prev = cur;
        cur = cur->next;
    }
}

inline bool HashTable::contains(const std::string& key) const {
    int h = hashString(key);
    Node* cur = buckets[h];
    while (cur != nullptr) {
        if (cur->key == key) return true;
        cur = cur->next;
    }
    return false;
}

inline int HashTable::size() const { return count; }

#endif