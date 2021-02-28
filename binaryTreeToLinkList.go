/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return
    }
    
    left, right := root.Left, root.Right
    
    root.Left = nil
    root.Right = left
    
    for root.Right != nil {
        root = root.Right
    }
    
    root.Right = right
    
    flatten(left)
    flatten(right)
}