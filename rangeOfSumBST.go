/**

Given the root node of a binary search tree, return the sum of values of all nodes with a value in the range [low, high].
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    
    var sum int
    if root.Val >= low && root.Val <= high {
      sum = root.Val
    }

    return sum + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
