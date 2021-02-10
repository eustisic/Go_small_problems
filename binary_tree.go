package main

import "fmt"

type node struct {
	data int
	leftChild *node
	rightChild *node
}

type BinaryTree struct {
	root *node
}

func (b *BinaryTree) initialize(data int) {
	b.root = &node{data: data, leftChild: nil, rightChild: nil} 
}

func (n *node) insert(data int) {
	if n == nil {
	  return
	}

	if n.data > data {
		if n.leftChild == nil {
			n.leftChild = &node{data: data, leftChild: nil, rightChild: nil}
		}
		n.leftChild.insert(data)

	} else if n.data < data {
		if n.rightChild == nil {
			n.rightChild = &node{data: data, leftChild: nil, rightChild: nil}
		}
		n.rightChild.insert(data)

	}
}

func (b *BinaryTree) printTree() {
  b.root.printAll()
}

func (n *node) printAll() {
  if n == nil {
    return
  }

  fmt.Printf("%d ", n.data)

  if n.leftChild != nil {
    n.leftChild.printAll()
  }

  if n.rightChild != nil {
    n.rightChild.printAll()
  }
}

func (n *node) traverse() {
	if n == nil {
		return
	}
	n.leftChild.traverse()

	fmt.Printf("%d ", n.data)

	n.rightChild.traverse()
}

func main() {

	tree := BinaryTree{}
	tree.initialize(50)
	tree.root.insert(25)
	tree.root.insert(75)
	tree.root.insert(33)
	tree.root.insert(18)
	tree.root.insert(88)
	tree.root.insert(66)
	tree.printTree()
}