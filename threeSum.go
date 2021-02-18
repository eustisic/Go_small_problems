func threeSum(nums []) [][]int {
	retArr := [][]int{{}}

	for i := 0; i < len(nums) - 2; i++ {
		for j := 1; j < len(nums) - 1; j++ {
			for k := 2; k < len(nums); k++ {
				num1 := nums[i]
				num2 := nums[j]
				num3 := nums[k]

				if num1 + num2 + num3 == 0 {
					retArr = append([]int{num1, num2, num3})
				}
			}
		}
	}
	return retArr
}