# Error Performance

This code benchmarks some way to assign an error

## How to test

```
cd $GOPATH/src/github.com/melodiez14/repokecil/_performance/exerror
go test -bench=. -benchmem
```

## Result

### Benchmark

```
BenchmarkNew-4          2000000000               0.46 ns/op            0 B/op          0 allocs/op
BenchmarkErrorf-4       10000000               124 ns/op              20 B/op          2 allocs/op
```

### Gcflags

```
./exerror_test.go:11:17: inlining call to errors.New
./exerror_test.go:9:22: BenchmarkNew b does not escape
./exerror_test.go:11:17: BenchmarkNew &errors.errorString literal does not escape
./exerror_test.go:11:17: BenchmarkNew error(&errors.errorString literal) does not escape
./exerror_test.go:15:25: BenchmarkErrorf b does not escape
```
