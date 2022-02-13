package main

import (
	"fmt"
	"math/rand"
)

const (
	maxVal = 1000
	pivot  = maxVal / 2
	arrLen = 100000
)

// randSlice generates slice of given len with random elements in [0-max]
func randSlice(len int, max int) []int {
	// create random array
	s := make([]int, len)
	for i := range s {
		s[i] = rand.Int() % (max + 1)
	}
	return s
}

//go:noinline
func sumLessThan(arr []int, threshold int) (sum int) {
	for _, elem := range arr {
		if elem > threshold {
			sum += elem % 10
		}
	}
	return sum
}

func main() {
	// create random array
	arr := randSlice(10000000, maxVal)

	// partition
	partition(arr, pivot)

	sum := sumLessThan(arr, pivot)
	fmt.Println(sum)

}
