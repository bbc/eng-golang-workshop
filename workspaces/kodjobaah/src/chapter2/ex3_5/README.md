# Benchmark for popcount


``` 
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: chapter2/ex3_5

BenchmarkPopCount-8                 	2000000000	         0.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountLoop-8             	100000000	        20.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountShift-8            	20000000	        60.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountClearRightmost-8   	30000000	        49.9 ns/op	       0 B/op	       0 allocs/op
```
