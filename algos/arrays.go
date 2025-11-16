package arrays

// RemoveDuplicatesFromSortedArrayInPlace - the idea is the following,
// we find duplicated values, then move all next value to the beginning of list (one position)
// at the end of list we store zero numbers
func RemoveDuplicatesFromSortedArrayInPlace(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nums
	}

	if len(nums) == 1 {
		return 1, nums
	}

	uniqueElements := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev > nums[i] {
			break // means that we have only zeros (deleted elements) in this position
		}

		if nums[i] != prev {
			prev = nums[i] // choose next unique value
			uniqueElements += 1
		} else {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1] // move all next elements one position
			}
			nums[len(nums)-1] = 0 // deleted duplicate goes to the end
			i -= 1                // we change i because array changed, so we have to check the same index again (in next step)
		}
	}
	return uniqueElements, nums[:uniqueElements]
}
