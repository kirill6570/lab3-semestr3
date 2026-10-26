#ifndef CBT_H
#define CBT_H

#include <string>
#include <iostream>
#include <queue>

class CompleteBinaryTree {
private:
    struct TreeNode {
        int data;
        TreeNode* left;
        TreeNode* right;
        TreeNode(int v) : data(v), left(nullptr), right(nullptr) {}
    };

    TreeNode* root;
    int count;

    void deleteTree(TreeNode* node);

public:
    CompleteBinaryTree() : root(nullptr), count(0) {}
    ~CompleteBinaryTree();

    CompleteBinaryTree(const CompleteBinaryTree&) = delete;
    CompleteBinaryTree& operator=(const CompleteBinaryTree&) = delete;

    void insert(int value);
    bool find(int value) const;
    bool isComplete() const;
    int size() const;
};

inline void CompleteBinaryTree::deleteTree(TreeNode* node) {
    if (!node) return;
    deleteTree(node->left);
    deleteTree(node->right);
    delete node;
}

inline CompleteBinaryTree::~CompleteBinaryTree() {
    deleteTree(root);
}

inline void CompleteBinaryTree::insert(int value) {
    TreeNode* newNode = new TreeNode(value);

    if (root == nullptr) {
        root = newNode;
        count++;
        return;
    }

    std::queue<TreeNode*> q;
    q.push(root);

    while (!q.empty()) {
        TreeNode* cur = q.front();
        q.pop();

        if (cur->left == nullptr) {
            cur->left = newNode;
            count++;
            return;
        } else {
            q.push(cur->left);
        }

        if (cur->right == nullptr) {
            cur->right = newNode;
            count++;
            return;
        } else {
            q.push(cur->right);
        }
    }
}

inline bool CompleteBinaryTree::find(int value) const {
    if (root == nullptr) return false;
    std::queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* cur = q.front();
        q.pop();
        if (cur->data == value) return true;
        if (cur->left) q.push(cur->left);
        if (cur->right) q.push(cur->right);
    }
    return false;
}

inline bool CompleteBinaryTree::isComplete() const {
    if (root == nullptr) return true;
    std::queue<TreeNode*> q;
    q.push(root);
    bool seenEmpty = false;

    while (!q.empty()) {
        TreeNode* cur = q.front();
        q.pop();

        if (cur->left) {
            if (seenEmpty) return false;
            q.push(cur->left);
        } else {
            seenEmpty = true;
        }

        if (cur->right) {
            if (seenEmpty) return false;
            q.push(cur->right);
        } else {
            seenEmpty = true;
        }
    }
    return true;
}

inline int CompleteBinaryTree::size() const { return count; }

#endif