func combinationSum(candidates []int, target int) (results [][]int) {
  var temp []int
  backtrack(candidates, target, temp, &results)
  return 
}

func backtrack(candidates []int, target int, temp []int, results *[][]int) {
  sum := sumOfArray(temp)
  if sum > target { return }
    
  if sum == target { // if the sum of the temp is equal to the target add to the 
    copyTemp := make([]int, len(temp))
    copy(copyTemp, temp)
    *results = append(*results, copyTemp)
    return
  }

    for i := 0; i < len(candidates); i++ {
      if len(temp) > 0 && candidates[i] < temp[len(temp)-1] {
          continue
      }
    
    temp = append(temp, candidates[i])     // take
    backtrack(candidates, target, temp, results) // explore
    temp = temp[:len(temp) - 1] // clean up
  }
}

func sumOfArray(arr []int) (sum int) {
    for _, v := range arr {
        sum += v
    }
    return
}
