package main

const times = 20

func main() {
	for i := 0; i < times; i++ {
		traverseColWise()
		traverseRowWise()
	}
}
