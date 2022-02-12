// +build !atomic

package main

func incrementManyTimes(val *int64, times int) {
	for i := 0; i < times; i++ {
		*val++
	}
}
