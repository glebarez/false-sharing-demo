// +build atomic

package main

import "sync/atomic"

func incrementManyTimes(addr *int64, times int) {
	for i := 0; i < times; i++ {
		atomic.AddInt64(addr, 1)
	}
}
