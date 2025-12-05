package miscellaneous

func CountPartitionsEvenSum(nums []int) int {
	evenCounter := 0
	for i := 0; i < len(nums)-1; i++ {
		leftArray := nums[:i+1]
		rightArray := nums[i+1:]
		sumLeft := 0
		sumRight := 0
		for _, item := range leftArray {
			sumLeft += item
		}
		for _, item := range rightArray {
			sumRight += item
		}

		if (sumLeft-sumRight)%2 == 0 {
			evenCounter++
		}
	}
	return evenCounter
}
