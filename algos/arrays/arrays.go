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
			nums[len(nums)-1] = -1 // deleted duplicate goes to the end
			i -= 1                 // we change i because array changed, so we have to check the same index again (in next step)
		}
	}
	return uniqueElements, nums[:uniqueElements]
}

func RemoveDuplicatesFromSortedArrayInPlaceAttemptTwo(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nums
	}

	if len(nums) == 1 {
		return 1, nums
	}

	lastUniqueIndex := 0
	uniqueElements := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			uniqueElements += 1
			if lastUniqueIndex+1 == i {
				lastUniqueIndex = i
			} else {
				nums[lastUniqueIndex+1] = nums[i]
				lastUniqueIndex += 1
			}
		}
		prev = nums[i]
	}
	return uniqueElements, nums[:uniqueElements]
}

func LongestCommonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0]
	}

	var prefixLen int = 0

labelOuterFor:
	for {
		var symbol uint8 = 0 // one character from string in Go is not string ! It is uint8 !
		for _, str := range words {
			if len(str) > prefixLen {
				if symbol == 0 {
					symbol = str[prefixLen]
				} else {
					if symbol != str[prefixLen] {
						break labelOuterFor
					}
				}
			} else {
				break labelOuterFor
			}
		}
		prefixLen += 1
	}

	if prefixLen == 0 {
		return ""
	}

	return words[0][:prefixLen]
}

func RemoveAllValElements(nums []int, val int) []int {
	lastIndex := len(nums) - 1
	equalToValNum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			break
		}

		if nums[i] == val {
			equalToValNum += 1
			nums[i] = nums[lastIndex]
			nums[lastIndex] = -1
			lastIndex--
			i--
		}
	}

	return nums[0 : len(nums)-equalToValNum]
}

func FindSingle(nums []int) int {
	mapElements := make(map[int]int)
	for _, num := range nums {
		_, exists := mapElements[num]
		if exists {
			delete(mapElements, num)
		} else {
			mapElements[num] = 1
		}
	}

	for k, v := range mapElements {
		if v == 1 {
			return k
		}
	}

	return -1
}
