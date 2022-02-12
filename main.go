package main

import (
	"sync"
)

func incrementParallel(val1, val2 *int64, times int) {
	// define waitgroup, so we can wait until background incrementing tasks are finished
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// run first incrementing task in background
	go func() {
		incrementManyTimes(val1, times)
		wg.Done()
	}()

	// run second incrementing task in background
	go func() {
		incrementManyTimes(val2, times)
		wg.Done()
	}()

	// wait for tasks to complete
	wg.Wait()
}

func main() {
	a := IntVars{}
	incrementParallel(&a.i1, &a.i2, 100000000)
}
