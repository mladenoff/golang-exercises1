package main

import "fmt"

var coins = [...]int{30, 14, 6, 2}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func makeChange(amount int, purse []int, minIdx int) []int {
	if amount == 0 {
		return purse
	}

	var bestPurse []int

	for idx, coin := range coins {
		if idx < minIdx {
			continue
		}

		if amount < coin {
			continue
		}

		tmpPurse := make([]int, len(purse), 2*len(purse))
		copy(tmpPurse, purse)
		tmpPurse = append(tmpPurse, coin)

		resultPurse := makeChange(amount-coin, tmpPurse, idx)
		if resultPurse == nil {
			continue
		}

		if bestPurse == nil {
			bestPurse = resultPurse
		} else if len(resultPurse) < len(bestPurse) {
			bestPurse = resultPurse
		}
	}

	return bestPurse
}

func main() {
	purse := makeChange(601, nil, 0)
	if purse != nil {
		fmt.Printf("%#v\n", purse)
	} else {
		fmt.Printf("COULD NOT MAKE CHANGE")
	}
}
