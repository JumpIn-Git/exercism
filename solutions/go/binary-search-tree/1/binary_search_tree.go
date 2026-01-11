package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{
		data: i,
	}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (node *BinarySearchTree) Insert(i int) {
	cur := node
	for {
		// Go left if smaller (also duplicates)
		if i <= cur.data {
			// If empty insert
			if cur.left == nil {
				cur.left = NewBst(i)
				break
			}
			// Move to the left
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = NewBst(i)
				break
			}
			cur = cur.right
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (node *BinarySearchTree) SortedData() (sorted []int) {
	if node != nil {
		sorted = append(sorted, node.left.SortedData()...)
		sorted = append(sorted, node.data)
		sorted = append(sorted, node.right.SortedData()...)
	}
	return sorted
}
