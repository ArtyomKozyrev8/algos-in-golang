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

func ClimbStairsRecursive(n int) int {
	storage := make(map[int]int)
	stepSizeOne := 1
	stepSizeTwo := 2

	var Inner func(x int) int

	Inner = func(x int) int {
		if x == 0 {
			return 1
		}
		if x < 0 {
			return 0
		}

		if _, exist := storage[x]; exist {
			return storage[x]
		} else {
			storage[x] = Inner(x-stepSizeOne) + Inner(x-stepSizeTwo)
			return storage[x]
		}
	}

	return Inner(n)
}
