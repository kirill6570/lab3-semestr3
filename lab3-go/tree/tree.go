package tree

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type CompleteBinaryTree struct {
	root *TreeNode
	size int
}

func New() *CompleteBinaryTree {
	return &CompleteBinaryTree{}
}

func (t *CompleteBinaryTree) Insert(value int) {
	newNode := &TreeNode{Data: value}

	if t.root == nil {
		t.root = newNode
		t.size++
		return
	}

	queue := []*TreeNode{t.root}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left == nil {
			cur.Left = newNode
			t.size++
			return
		} else {
			queue = append(queue, cur.Left)
		}

		if cur.Right == nil {
			cur.Right = newNode
			t.size++
			return
		} else {
			queue = append(queue, cur.Right)
		}
	}
}

func (t *CompleteBinaryTree) Find(value int) bool {
	if t.root == nil {
		return false
	}

	queue := []*TreeNode{t.root}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Data == value {
			return true
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return false
}

func (t *CompleteBinaryTree) IsComplete() bool {
	if t.root == nil {
		return true
	}

	queue := []*TreeNode{t.root}
	seenEmpty := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left != nil {
			if seenEmpty {
				return false
			}
			queue = append(queue, cur.Left)
		} else {
			seenEmpty = true
		}

		if cur.Right != nil {
			if seenEmpty {
				return false
			}
			queue = append(queue, cur.Right)
		} else {
			seenEmpty = true
		}
	}
	return true
}

func (t *CompleteBinaryTree) Size() int {
	return t.size
}

func (t *CompleteBinaryTree) PrintByLevels() {
	if t.root == nil {
		fmt.Println("дерево пусто")
		return
	}

	queue := []*TreeNode{t.root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Print(cur.Data, " ")
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		fmt.Println()
	}
}