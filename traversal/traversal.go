package main

const (
	rows = 2000
	cols = 3000
)

var arr [rows][cols]int64

//go:noinline
func traverseColWise() {
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			arr[r][c] = 1
		}
	}
}

//go:noinline
func traverseRowWise() {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			arr[r][c] = 1
		}
	}
}
