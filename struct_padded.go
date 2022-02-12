// +build padded

package main

import "golang.org/x/sys/cpu"

type IntVars struct {
	i1 int64
	_  cpu.CacheLinePad // padding
	i2 int64
}
