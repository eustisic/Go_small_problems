/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    queue := []*TreeNode{}
    result := [][]int{}
    queue = append(queue, root)
    
    for len(queue) != 0 {
        levels := []int{}
        div := len(queue)
        
        for i := 0; i < div; i++ {
            r := queue[0]
            
            if r.Right != nil {
                queue = append(queue, r.Right)
            }
            
            if r.Left != nil {
                queue = append(queue, r.Left)
            }
            
            levels = append(levels, r.Val)
        }
    
        result = append(result, levels)
        queue = queue[1:]
    }
    
    return result
    
}

/*

create a queue
add roo to queue
create a result nexted array

while the queue is not empty

initialize an array levels
set queue length to variable div

from 0 to div
add queue[0] to levels
add queue[0] children to queue
dequeue queue[0]

append levels to result


return result
*/

