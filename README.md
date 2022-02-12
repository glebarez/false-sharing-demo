### Example of CPU cache false-sharing in Go.
A simple example where 2 integer variables are incremented concurrently.<br>
Baseline version suffers from false-sharing due to values share same cache line:
```go
type IntVars struct {
	i1 int64
	i2 int64
}
```

Optimized version eliminates false sharing by introducing padding between memory locations:
```go
type IntVars struct {
	i1 int64
	_  cpu.CacheLinePad // padding
	i2 int64
}
```


### Reproducing results
1. Install benchstat:
```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

2. Run benchmarks for simple increments ```a++```:
```bash
make bench

name                          old time/op  new time/op  delta
Increment1Value-2             1.66ns ± 8%  1.71ns ± 7%     ~     (p=0.421 n=5+5)
Increment2ValuesInParallel-2  2.34ns ± 5%  1.59ns ± 3%  -32.23%  (p=0.008 n=5+5)
```

3. Run benchmarks for atomic increments: ```atomic.AddInt64(addr, 1)```:
```bash
make bench-atomic

name                          old time/op  new time/op  delta
Increment1Value-2             5.65ns ± 5%  5.85ns ± 6%     ~     (p=0.310 n=5+5)
Increment2ValuesInParallel-2  41.6ns ±10%   5.4ns ± 8%  -87.12%  (p=0.008 n=5+5)
```

### CPU cache miss stats
On Linux, one can measure L1 cache misses to demonstrate false-sharing.
1. Build executables for both original and optimized versions:
```bash
make build
```

2. Run perf for both executables and compare numbers:
```bash
perf stat -B -e L1-dcache-load-misses ./test

 Performance counter stats for './test':

         8,954,010      L1-dcache-load-misses
```

```bash
perf stat -B -e L1-dcache-load-misses ./test-padded

 Performance counter stats for './test-padded':

           204,287      L1-dcache-load-misses
```