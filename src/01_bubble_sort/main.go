package main

import (
	"fmt"
	"math/rand"
)

// bubbleSort
func bubbleSort(nums []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				swap(&nums[i], &nums[i+1])
				sorted = false
			} else {
				continue
			}
		}
	}
	return nums
}

// randomInts
func randomInts() []int {
	rand.Seed(825)
	randoms := []int{}
	for i := 0; i < 5; i++ {
		randoms = append(randoms, rand.Intn(100))
	}
	return randoms
}

// swap
func swap(xp, yp *int) {
	xval := *xp
	yval := *yp
	*xp, *yp = yval, xval
}

func main() {
	unsortedArray := randomInts()
	fmt.Printf("Starting slice: %v\n", unsortedArray)
	unsortedArray = bubbleSort(unsortedArray)
	fmt.Printf("Ending slice: %v\n", unsortedArray)
}
