package main

import (
	"testing"
)

func BenchmarkArrayProcessing(b *testing.B) {
	arr := randSlice(arrLen, maxVal)

	partition(arr, pivot)

	// reset timer
	b.ResetTimer()

	// process
	for n := 0; n < b.N; n++ {
		sumLessThan(arr, pivot)
	}
}
