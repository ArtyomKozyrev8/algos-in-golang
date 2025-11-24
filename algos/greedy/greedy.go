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
