// +build cols

package main

import "testing"

func BenchmarkTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		traverseColWise()
	}
}
