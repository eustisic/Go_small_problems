package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func addNode(n *TreeNode, v int) {
	if v < n.Val {

		if n.Left == nil {
			n.Left = &TreeNode{Val: v}
		} else {
			addNode(n.Left, v)
		}
	}

	if n.Val <= v {

		if n.Right == nil {
			n.Right = &TreeNode{Val: v}
		} else {
			addNode(n.Right, v)
		}
	}
}

func createTree(arr []int) *TreeNode {
	var root *TreeNode

	for i, v := range arr {
		if i == 0 {
			root = &TreeNode{Val: v}
		} else {
			addNode(root, v)
		}
	}
	return root
}

func printAll(n *TreeNode) {
	fmt.Printf("%d ", n.Val)

	if n.Left != nil {
		printAll(n.Left)
	}

	if n.Right != nil {
		printAll(n.Right)
	}
}

func maxDepth(n *TreeNode, d int, md *int) {
	if n == nil {
		if d - 1 > *md || *md == 0 {
		  *md = d - 1
		}
		return
	}

	
	maxDepth(n.Left, d + 1, md)

	maxDepth(n.Right, d + 1, md)
	
}

func minDepth(n *TreeNode, d int, md *int) {
	if n == nil {
		if d - 1 < *md || *md == 0 {
		  *md = d - 1
		}
		return
	}

	
	minDepth(n.Left, d + 1, md)

	minDepth(n.Right, d + 1, md)
	
}

func main() {
	t1 := createTree([]int{9, 3, 20, 15, 17})
	printAll(t1)
	var max int
	var min int
	maxDepth(t1, 1, &max)
	minDepth(t1, 1, &min)
	fmt.Println()
	fmt.Printf("Min Depth -> %d \n", min)
	fmt.Printf("Max Depth -> %d \n", max)
}