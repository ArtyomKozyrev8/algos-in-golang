package greedy

import (
	"testing"
)

type BuyChocoTest struct {
	prices []int
	money  int
	expect int
}

func TestBuyChoco(t *testing.T) {
	tests := []BuyChocoTest{
		{prices: []int{1, 2, 3}, money: 2, expect: 2},
		{prices: []int{4, 1, 3, 2, 5}, money: 2, expect: 2},
		{prices: []int{4, 1, 3, 2, 5}, money: 3, expect: 0},
		{prices: []int{}, money: 5, expect: 5},
		{prices: []int{1}, money: 5, expect: 5},
		{prices: []int{5, 4, 3, 2, 1}, money: 3, expect: 0},
	}

	for _, test := range tests {
		res := BuyChoco(test.prices, test.money)
		if res != test.expect {
			t.Errorf("%v != %v", res, test.expect)
		}
	}
}
