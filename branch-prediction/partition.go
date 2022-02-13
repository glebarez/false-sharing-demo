// +build !rand

package main

func partition(arr []int, pivot int) {
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	low, high := 0, len(arr)-1

	// Index of smaller element and indicates the
	// right position of pivot found so far
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++ // increment index of smaller element
			swap(i, j)
		}
	}
	swap(i+1, high)
}
