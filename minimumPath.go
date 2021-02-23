


/*
BRUTE FORCE 
calculate all paths, store in a hash and minimize the hash

DYNAMIC APPROACH
bottom up recursive algorithm

calculate the simplist path and compare the values of each

enter the smaller of the values into the next recursive call
*/

func minPathSum(grid [][]int) int {
    cache := map[[2]int]int{}
    return helper(0, 0, grid, cache)
}

func helper(row int, col int, grid [][]int, cache map[[2]int]int) int {
    if row + 1 == len(grid) && col + 1 == len(grid[0]) {
        return grid[row][col]
    }

    key := [2]int{row, col}

    _, ok := cache[key]
    
    if row+1 < len(grid) && col+1 < len(grid[0]) {
        
        if !ok {
            cache[key] = grid[row][col] + min(helper(row+1, col, grid, cache), helper(row, col+1, grid, cache))
        }
        
    } else if row+1 < len(grid) {
        
        if !ok {
            cache[key] = grid[row][col] + helper(row+1, col, grid, cache)
        }
    } else if col+1 < len(grid[0]) {
        
        if !ok {
            cache[key] = grid[row][col] + helper(row, col+1, grid, cache)
        }
    }

    return cache[key]
}

func min(val1, val2 int) int {
    if val1 < val2 {
        return val1
    }
    return val2
}



