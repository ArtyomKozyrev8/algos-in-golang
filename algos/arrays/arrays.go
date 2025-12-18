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

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var digits []int

	for x >= 10 {
		digit := x % 10
		digits = append(digits, digit)
		x = (x - digit) / 10
	}

	if x != 0 {
		digits = append(digits, x)
	}

	for i := 0; i < len(digits); i++ {
		if i == len(digits)-1 {
			return true
		}

		if i == (len(digits)-len(digits)%2)+1 {
			return true
		}

		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}

func SearchInsertIndexReturn(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target < nums[0] {
		return 0
	}

	newNums := nums
	globalIndexLeft := 0
	for {
		if len(newNums) == 1 {
			if newNums[0] == target {
				return globalIndexLeft
			} else if newNums[0] < target {
				return globalIndexLeft + 1
			} else {
				return globalIndexLeft
			}
		}

		middleIndex := (len(newNums) - len(newNums)%2) / 2
		if middleIndex == 0 {
			return globalIndexLeft
		}
		if nums[middleIndex] == target {
			return globalIndexLeft + middleIndex
		} else if newNums[middleIndex] > target {
			newNums = newNums[:middleIndex]
		} else {
			newNums = newNums[middleIndex:]
			globalIndexLeft += middleIndex
		}
	}
}

func CoinChangeBreadthFirst(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	amounts := []int{amount}
	minNumberOfCoins := -1

	storage := make(map[int]bool) // key - restAmount

	coinsNumber := 0
	for len(amounts) > 0 {
		var newAmounts []int
		for _, amountLocal := range amounts {
			for _, coin := range coins {
				restAmount := amountLocal - coin

				if restAmount != 0 {
					if _, exists := storage[restAmount]; exists {
						continue
					}
				}

				if restAmount == 0 {
					return coinsNumber + 1
				}

				if restAmount > 0 {
					newAmounts = append(newAmounts, restAmount)
					storage[restAmount] = true
				}
			}
		}
		amounts = newAmounts
		coinsNumber += 1
	}
	return minNumberOfCoins
}

func CoinChangeRecursionAttemptOne(coins []int, amount int) int {
	var innerFunc func(amountL int, usedCoinsNumber int)

	seenWays := make(map[int]int)
	minNumberOfCoins := -1

	innerFunc = func(amountL int, usedCoinsNumberL int) {
		if amountL == 0 {
			if minNumberOfCoins == -1 {
				minNumberOfCoins = usedCoinsNumberL
			} else {
				if usedCoinsNumberL < minNumberOfCoins {
					minNumberOfCoins = usedCoinsNumberL
				}
			}
		}

		if amountL < 0 {
			return
		}

		for _, coin := range coins {
			val, exist := seenWays[amountL-coin]
			if val > usedCoinsNumberL+1 || !exist {
				if amountL-coin >= 0 {
					seenWays[amountL-coin] = usedCoinsNumberL + 1
					innerFunc(amountL-coin, usedCoinsNumberL+1)
				}
			}
		}
	}

	innerFunc(amount, 0)
	return minNumberOfCoins
}
