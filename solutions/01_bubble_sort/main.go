package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(nums []int) {
	didSwap := true
	for didSwap {
		didSwap = false
		for idx := range nums {
			if idx == len(nums)-1 {
				continue
			}

			if nums[idx+1] < nums[idx] {
				// nums[idx+1], nums[idx] = nums[idx], nums[idx+1]
				swap(&nums[idx], &nums[idx+1])
				didSwap = true
			}
		}
	}
}

func randomInts(targetLength int) []int {
	nums := []int{}
	for len(nums) < targetLength {
		num := rand.Intn(1000)
		nums = append(nums, num)
	}

	return nums
}

func swap(xp *int, yp *int) {
	*yp, *xp = *xp, *yp
}

func main() {
	nums := randomInts(20)
	bubbleSort(nums)

	x, y := 123, 456
	swap(&x, &y)
	fmt.Printf("%v|%v\n", x, y)
	fmt.Println(nums)
}
