package main

import (
	"testing"
)

func BenchmarkIncrement1Value(b *testing.B) {
	a := new(IntVars)
	incrementManyTimes(&a.i1, b.N)
}

func BenchmarkIncrement2ValuesInParallel(b *testing.B) {
	a := new(IntVars)
	incrementParallel(&a.i1, &a.i2, b.N)
}
