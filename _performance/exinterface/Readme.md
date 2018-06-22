# Interface Performance

This code benchmarks direct function vs interface call

## How to test

```
cd $GOPATH/src/github.com/melodiez14/repokecil/_performance/exinterface
go test -bench=. -benchmem
```

## Result

### Benchmark

```
BenchmarkUser_Scan-4                    2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkUser_ScanInterface1-4          200000000                8.82 ns/op            0 B/op          0 allocs/op
BenchmarkUser_ScanInterface2-4          300000000                5.33 ns/op            0 B/op          0 allocs/op
```

### Gcflags

```
./exinterface.go:9:6: can inline New
./exinterface.go:16:6: can inline (*User).Print
./exinterface.go:20:6: can inline (*User).Scan
./exinterface_test.go:6:10: inlining call to New
./exinterface_test.go:8:9: inlining call to (*User).Scan
./exinterface_test.go:13:10: inlining call to New
./exinterface_test.go:20:10: inlining call to New
./exinterface.go:16:10: (*User).Print u does not escape
./exinterface.go:20:41: leaking param: firstname
./exinterface.go:20:41: leaking param: lastname
./exinterface.go:20:41: (*User).Scan u does not escape
./exinterface_test.go:5:28: BenchmarkUser_Scan b does not escape
./exinterface_test.go:8:4: BenchmarkUser_Scan u does not escape
./exinterface_test.go:15:17: &u escapes to heap
./exinterface_test.go:15:17: &u escapes to heap
./exinterface_test.go:13:2: moved to heap: u
./exinterface_test.go:12:38: BenchmarkUser_ScanInterface1 b does not escape
./exinterface_test.go:22:16: IPrinter(&u) escapes to heap
./exinterface_test.go:22:17: &u escapes to heap
./exinterface_test.go:20:2: moved to heap: u
./exinterface_test.go:19:38: BenchmarkUser_ScanInterface2 b does not escape
```

##### Explanation

###### Tips: Before read this explanation. You need to understand how do the stack and heap memory works and GO escape analysis

1.  `./exinterface_test.go:8:4: BenchmarkUser_Scan u does not escape`
    This is the direct function call. The u object wont be escaped to the heap and keep in stack memory
2.  `./exinterface_test.go:15:17: &u escapes to heap`
    This is use an interface call `IPrinter.Scan(&u, "Risal", "Falah")`. It cause `u` will be escaped to the heap.
3.  `./exinterface_test.go:22:16: IPrinter(&u) escapes to heap`
    This is use an interface call `IPrinter.Scan(&u, "Risal", "Falah")`. It cause `u` will be escaped to the heap.
