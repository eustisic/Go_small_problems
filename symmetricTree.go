/**

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center)

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return checkInvert(root.Left, root.Right)
}


func checkInvert(l, r *TreeNode) bool {

    if l == nil && r == nil {
        
      return true
        
    } else if l != nil && r != nil {
        
        if l.Val != r.Val {
            return false
        }
    
        return checkInvert(l.Left, r.Right) && checkInvert(l.Right, r.Left)
    }
    
    return false
}
