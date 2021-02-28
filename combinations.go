func combine(n int, k int) [][]int {
  var results [][]int
  var temp []int
  backtrack(1, n, k, temp, &results)
  return results 
}

func backtrack(s int, n int, k int, temp []int, results *[][]int) {
    if len(temp) > k { return }
    
  if len(temp) == k { // also can return if terminal condition (lead node)
    copyTemp := make([]int, len(temp))
    copy(copyTemp, temp)
    *results = append(*results, copyTemp)
    return
  }

  for i := s; i <= n; i++ {
      if len(temp) > 0 && i <= temp[len(temp)-1] {
          continue
      }
    
    temp = append(temp, i)     // take
    backtrack(s+1, n, k, temp, results) // explore
    temp = temp[:len(temp) - 1] // clean up
  }
}
