package greedy

func BuyChoco(prices []int, money int) int {
	if len(prices) < 1 {
		return money // nothing to buy
	}

	firstSmallest := -1
	secondSmallest := -1

	for _, price := range prices {
		if firstSmallest == -1 {
			firstSmallest = price
		} else if secondSmallest == -1 {
			if firstSmallest > price {
				secondSmallest = firstSmallest
				firstSmallest = price
			} else {
				secondSmallest = price
			}
		} else {
			if price <= firstSmallest {
				secondSmallest = firstSmallest
				firstSmallest = price
			} else if price <= secondSmallest {
				secondSmallest = price
			}
		}
	}

	if secondSmallest == -1 {
		return money // we need to buy exactly two chocolates
	}

	moneyLeft := money - firstSmallest - secondSmallest
	if moneyLeft >= 0 {
		return moneyLeft
	} else {
		return money
	}
}

func LongestPalindrome(s string) int {
	result := 0
	letters := make(map[int32]int)
	for _, letter := range s {
		val, exist := letters[letter]
		if exist == false {
			letters[letter] = 1
		} else if val == 1 {
			letters[letter] = 0
			result += 2
			delete(letters, letter)
		} else {
			letters[letter] = 1
		}
	}
	if len(letters) > 0 {
		result += 1
	}

	return result
}
