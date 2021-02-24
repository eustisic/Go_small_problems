


/*
BRUTE FORCE 
calculate all paths, store in a hash and minimize the hash

DYNAMIC APPROACH
bottom up recursive algorithm

calculate the simplist path and compare the values of each

enter the smaller of the values into the next recursive call
*/

1 1 3      1 2 _

4 5 1      5 _ _

7 8 1      _ _ _

func minPathSum(grid [][]int) int {

  cache := map[[2]int]int{}


  for row := 0; row < len(grid); row++ {
    for col := 0; col < len(grid[0]); col++ {

      key := [2]int{row, col}
      loc := grid[row][col]

      if row == 0 && col == 0 {
        cache[key] = loc
      } else if row == 0 {
        cache[key] = loc + cache[[2]int{row, col-1}]
      } else if col == 0 {
        cache[key] = loc + cache[[2]int{row-1, col}]
      } else {
        cache[key] = loc + min(cache[[2]int{row-1, col}], cache[[2]int{row, col-1}])
      }

    }
  }

    return cache[[2]int{len(grid) - 1, len(grid[0]) - 1}]
}

func min(val1, val2 int) int {
  if val1 < val2 {
    return val1
  }
  return val2
}

  // if row + 1 == len(grid) && col + 1 == len(grid[0]) {
  // 	return grid[row][col]
  // }

//   key := [2]int{row, col}

//   _, ok := cache[key]

//   if !ok {
//     if row+1 < len(grid) && col+1 < len(grid[0]) {

//       cache[key] = grid[row][col] + min(helper(row+1, col, grid, cache), helper(row, col+1, grid, cache))
      
//     } else if row+1 < len(grid) {

//     	cache[key] = grid[row][col] + helper(row+1, col, grid, cache)

//     } else if col+1 < len(grid[0]) {

//     	cache[key] = grid[row][col] + helper(row, col+1, grid, cache)
//     }
//   }
  
//   return cache[[2]int{0, 0}]
// }





