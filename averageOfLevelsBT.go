func averageOfLevels(root *TreeNode) []float64 {
    queue := []*TreeNode{}
    result := []float64{}
    queue = append(queue, root)
    
    for len(queue) != 0 {
        level := 0
        div := len(queue)
        
        for i := 0; i < div; i++ {
            r := queue[0]
            
            if r.Left != nil {
                queue = append(queue, r.Left)
            }
            if r.Right != nil {
                queue = append(queue, r.Right)
            }

            level += r.Val
            queue = queue[1:]
        }
        
        
        result = append(result, float64(level) / float64(div))
    }
    
    return result
}