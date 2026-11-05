#include <gtest/gtest.h>

#include "DynamicArray.h"
#include "SinglyList.h"
#include "DoublyList.h"
#include "Stack.h"
#include "Queue.h"
#include "HashTable.h"
#include "CompleteBinaryTree.h"
#include "Serializer.h"
#include "TextSerializer.h"

// ===== DynamicArray =====
TEST(DynamicArrayTest, InitiallyEmpty) {
    DynamicArray arr;
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(DynamicArrayTest, PushBack) {
    DynamicArray arr;
    arr.pushBack("10");
    arr.pushBack("20");
    EXPECT_EQ(arr.getSize(), 2);
    EXPECT_EQ(arr.get(0), "10");
    EXPECT_EQ(arr.get(1), "20");
}

TEST(DynamicArrayTest, Set) {
    DynamicArray arr;
    arr.pushBack("10");
    arr.set(0, "99");
    EXPECT_EQ(arr.get(0), "99");
}

TEST(DynamicArrayTest, RemoveAt) {
    DynamicArray arr;
    arr.pushBack("10");
    arr.pushBack("20");
    arr.pushBack("30");
    arr.removeAt(1);
    EXPECT_EQ(arr.getSize(), 2);
    EXPECT_EQ(arr.get(1), "30");
}

TEST(DynamicArrayTest, PopBack) {
    DynamicArray arr;
    arr.pushBack("10");
    arr.popBack();
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(DynamicArrayTest, OutOfBounds) {
    DynamicArray arr;
    EXPECT_EQ(arr.get(-1), "");
    EXPECT_EQ(arr.get(100), "");
}
TEST(DynamicArrayTest, MoveSemantics) {
    DynamicArray arr1;
    arr1.pushBack("hello");
    arr1.pushBack("world");
    
    DynamicArray arr2(std::move(arr1));
    EXPECT_EQ(arr2.getSize(), 2);
    EXPECT_EQ(arr1.getSize(), 0); 
    
    DynamicArray arr3;
    arr3 = std::move(arr2);
    EXPECT_EQ(arr3.getSize(), 2);
    EXPECT_EQ(arr2.getSize(), 0);
}
TEST(DynamicArrayTest, SetRemoveAndBounds) {
    DynamicArray arr;
    arr.pushBack("test");
    arr.set(5, "fail");
    arr.removeAt(5);
    arr.get(5);
    EXPECT_EQ(arr.getSize(), 1); 
    arr.set(0, "new_test");
    EXPECT_EQ(arr.get(0), "new_test");
    
    arr.removeAt(0);
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(DynamicArrayTest, MoveAssignmentAndEnsureCapacity) {
    DynamicArray arr1;
    for (int i = 0; i < 10; ++i) {
        arr1.pushBack("item");
    }
    EXPECT_EQ(arr1.getSize(), 10); 
    DynamicArray arr2;
    arr2 = std::move(arr1);
    EXPECT_EQ(arr2.getSize(), 10);
    EXPECT_EQ(arr1.getSize(), 0); 
}

TEST(DynamicArrayTest, PrintCoverage) {
    DynamicArray arr;
    arr.print(); 
    
    arr.pushBack("A");
    arr.pushBack("B");
    arr.print();
    
    EXPECT_EQ(arr.getSize(), 2);
}

// ===== SinglyList =====
TEST(SinglyListTest, InitiallyEmpty) {
    SinglyList list;
    EXPECT_EQ(list.getSize(), 0);
}

TEST(SinglyListTest, PushFront) {
    SinglyList list;
    list.pushFront("a");
    list.pushFront("b");
    EXPECT_EQ(list.getSize(), 2);
    EXPECT_TRUE(list.find("a"));
    EXPECT_TRUE(list.find("b"));
}

TEST(SinglyListTest, PushBack) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    EXPECT_EQ(list.getSize(), 2);
}

TEST(SinglyListTest, PopFront) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.popFront();
    EXPECT_EQ(list.getSize(), 1);
    EXPECT_FALSE(list.find("a"));
    EXPECT_TRUE(list.find("b"));
}

TEST(SinglyListTest, PopBack) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.popBack();
    EXPECT_EQ(list.getSize(), 1);
    EXPECT_FALSE(list.find("b"));
}

TEST(SinglyListTest, RemoveValue) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.pushBack("c");
    list.remove("b");
    EXPECT_EQ(list.getSize(), 2);
    EXPECT_FALSE(list.find("b"));
}

TEST(SinglyListTest, FindMissing) {
    SinglyList list;
    list.pushBack("a");
    EXPECT_FALSE(list.find("z"));
}

// ===== DoublyList =====
TEST(DoublyListTest, InitiallyEmpty) {
    DoublyList list;
    EXPECT_EQ(list.getSize(), 0);
}

TEST(DoublyListTest, PushFrontBack) {
    DoublyList list;
    list.pushBack("b");
    list.pushFront("a");
    list.pushBack("c");
    EXPECT_EQ(list.getSize(), 3);
    EXPECT_TRUE(list.find("a"));
    EXPECT_TRUE(list.find("c"));
}

TEST(DoublyListTest, PopFront) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.popFront();
    EXPECT_EQ(list.getSize(), 1);
    EXPECT_FALSE(list.find("a"));
}

TEST(DoublyListTest, PopBack) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.popBack();
    EXPECT_EQ(list.getSize(), 1);
    EXPECT_FALSE(list.find("b"));
}

TEST(DoublyListTest, RemoveMiddle) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.pushBack("c");
    list.remove("b");
    EXPECT_EQ(list.getSize(), 2);
    EXPECT_FALSE(list.find("b"));
    EXPECT_TRUE(list.find("a"));
    EXPECT_TRUE(list.find("c"));
}

// ===== Stack =====
TEST(StackTest, InitiallyEmpty) {
    Stack s;
    EXPECT_TRUE(s.isEmpty());
    EXPECT_EQ(s.size(), 0);
}

TEST(StackTest, PushPeek) {
    Stack s;
    s.push("10");
    s.push("20");
    EXPECT_EQ(s.size(), 2);
    EXPECT_EQ(s.peek(), "20");
}

TEST(StackTest, Pop) {
    Stack s;
    s.push("10");
    s.push("20");
    s.pop();
    EXPECT_EQ(s.size(), 1);
    EXPECT_EQ(s.peek(), "10");
}

TEST(StackTest, PopEmpty) {
    Stack s;
    s.pop();
    EXPECT_TRUE(s.isEmpty());
}

// ===== Queue =====
TEST(QueueTest, InitiallyEmpty) {
    Queue q;
    EXPECT_TRUE(q.isEmpty());
    EXPECT_EQ(q.size(), 0);
}

TEST(QueueTest, PushFront) {
    Queue q;
    q.push("10");
    q.push("20");
    EXPECT_EQ(q.size(), 2);
    EXPECT_EQ(q.front(), "10");
}

TEST(QueueTest, Pop) {
    Queue q;
    q.push("10");
    q.push("20");
    q.pop();
    EXPECT_EQ(q.size(), 1);
    EXPECT_EQ(q.front(), "20");
}

TEST(QueueTest, PopEmpty) {
    Queue q;
    q.pop();
    EXPECT_TRUE(q.isEmpty());
}

// ===== HashTable =====
TEST(HashTableTest, InitiallyEmpty) {
    HashTable ht;
    EXPECT_EQ(ht.size(), 0);
    EXPECT_FALSE(ht.contains("apple"));
}

TEST(HashTableTest, AddContains) {
    HashTable ht;
    ht.add("apple");
    ht.add("banana");
    EXPECT_EQ(ht.size(), 2);
    EXPECT_TRUE(ht.contains("apple"));
    EXPECT_TRUE(ht.contains("banana"));
}

TEST(HashTableTest, Duplicate) {
    HashTable ht;
    ht.add("apple");
    ht.add("apple");
    EXPECT_EQ(ht.size(), 1);
}

TEST(HashTableTest, Remove) {
    HashTable ht;
    ht.add("apple");
    ht.add("banana");
    ht.remove("apple");
    EXPECT_EQ(ht.size(), 1);
    EXPECT_FALSE(ht.contains("apple"));
    EXPECT_TRUE(ht.contains("banana"));
}

// ===== CompleteBinaryTree =====
TEST(CBTTest, InitiallyEmpty) {
    CompleteBinaryTree tree;
    EXPECT_EQ(tree.size(), 0);
    EXPECT_TRUE(tree.isComplete());
}

TEST(CBTTest, Insert) {
    CompleteBinaryTree tree;
    tree.insert(10);
    tree.insert(20);
    tree.insert(30);
    EXPECT_EQ(tree.size(), 3);
    EXPECT_TRUE(tree.isComplete());
}

TEST(CBTTest, Find) {
    CompleteBinaryTree tree;
    tree.insert(10);
    tree.insert(20);
    tree.insert(30);
    tree.insert(40);
    EXPECT_TRUE(tree.find(20));
    EXPECT_TRUE(tree.find(40));
    EXPECT_FALSE(tree.find(99));
}

TEST(CBTTest, IsCompleteAfterInserts) {
    CompleteBinaryTree tree;
    for (int i = 1; i <= 7; ++i) tree.insert(i);
    EXPECT_TRUE(tree.isComplete());
    EXPECT_EQ(tree.size(), 7);
}

// ===== Сериализация =====
TEST(SerializationTest, BinaryDynamicArray) {
    DynamicArray arr;
    arr.pushBack("10");
    arr.pushBack("20");
    arr.pushBack("30");

    Serializer::saveBinary(arr, "test_arr.bin");
    DynamicArray loaded = Serializer::loadBinary("test_arr.bin");

    EXPECT_EQ(loaded.getSize(), 3);
    EXPECT_EQ(loaded.get(0), "10");
    EXPECT_EQ(loaded.get(1), "20");
    EXPECT_EQ(loaded.get(2), "30");
}

TEST(SerializationTest, TextSinglyList) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.pushBack("c");

    TextSerializer::saveText(list, "test_list.txt");
    SinglyList loaded = TextSerializer::loadText("test_list.txt");

    EXPECT_EQ(loaded.getSize(), 3);
    EXPECT_TRUE(loaded.find("a"));
    EXPECT_TRUE(loaded.find("c"));
}