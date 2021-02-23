/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
*/


func minimumTotal(triangle [][]int) int {
    cache := map[[2]int]int{}
    return helper(0, 0, triangle, cache)
}

func helper(row, col int, triangle [][]int, cache map[[2]int]int) int {
    if row + 1 == len(triangle) {
        return triangle[row][col]
    }
    
    key := [2]int{row, col}
    // check if value is cached
    _, ok := cache[key]
    
    if !ok {
        cache[key] = triangle[row][col] + min(helper(row+1, col, triangle, cache), helper(row+1, col+1, triangle, cache))
    }
    
    return cache[key]
}

func min(val1, val2 int) int {
    if val1 < val2 {
        return val1
    }
    return val2
}