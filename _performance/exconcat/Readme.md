# Cocat Performance

This code benchmarks some way to concat string

## How to test

```
cd $GOPATH/src/github.com/melodiez14/repokecil/_performance/exconcat
go test -bench=. -benchmem
```

## Result

### Benchmark

```
BenchmarkSprintfShortString-4           10000000               137 ns/op              32 B/op          1 allocs/op
BenchmarkPlusShortString-4              2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkJoinShortString-4              20000000                71.5 ns/op            32 B/op          1 allocs/op
BenchmarkBuilderShortString-4           20000000               103 ns/op              48 B/op          2 allocs/op
BenchmarkBufferShortString-4            20000000                92.7 ns/op           112 B/op          1 allocs/op
BenchmarkSprintfLongString-4             5000000               270 ns/op             576 B/op          1 allocs/op
BenchmarkPlusLongString-4               2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkJoinLongString-4               10000000               206 ns/op             576 B/op          1 allocs/op
BenchmarkBuilderLongString-4             5000000               316 ns/op             864 B/op          2 allocs/op
BenchmarkBufferLongString-4              2000000               770 ns/op            1872 B/op          4 allocs/op
BenchmarkSprintfStringInt64_1-4         10000000               142 ns/op              32 B/op          1 allocs/op
BenchmarkSprintfStringInt64_2-4         10000000               158 ns/op              32 B/op          1 allocs/op
BenchmarkPlusStringInt64-4              20000000                63.4 ns/op             4 B/op          1 allocs/op
```

### Gcflags

```
./exconcat_test.go:50:21: inlining call to strings.(*Builder).String
./exconcat_test.go:63:20: inlining call to bytes.(*Buffer).String
./exconcat_test.go:106:21: inlining call to strings.(*Builder).String
./exconcat_test.go:119:20: inlining call to bytes.(*Buffer).String
./exconcat_test.go:17:19: str escapes to heap
./exconcat_test.go:17:19: id escapes to heap
./exconcat_test.go:11:37: BenchmarkSprintfShortString b does not escape
./exconcat_test.go:17:18: BenchmarkSprintfShortString ... argument does not escape
./exconcat_test.go:21:34: BenchmarkPlusShortString b does not escape
./exconcat_test.go:31:34: BenchmarkJoinShortString b does not escape
./exconcat_test.go:37:28: BenchmarkJoinShortString []string literal does not escape
./exconcat_test.go:41:37: BenchmarkBuilderShortString b does not escape
./exconcat_test.go:48:10: BenchmarkBuilderShortString builder does not escape
./exconcat_test.go:49:10: BenchmarkBuilderShortString builder does not escape
./exconcat_test.go:50:14: BenchmarkBuilderShortString builder does not escape
./exconcat_test.go:50:21: BenchmarkBuilderShortString &strings.b·2.buf does not escape
./exconcat_test.go:61:9: buffer escapes to heap
./exconcat_test.go:60:7: moved to heap: buffer
./exconcat_test.go:62:9: buffer escapes to heap
./exconcat_test.go:54:36: BenchmarkBufferShortString b does not escape
./exconcat_test.go:63:13: BenchmarkBufferShortString buffer does not escape
./exconcat_test.go:63:20: BenchmarkBufferShortString string(bytes.b·2.buf[bytes.b·2.off:]) does not escape
./exconcat_test.go:73:19: strone escapes to heap
./exconcat_test.go:73:19: strtwo escapes to heap
./exconcat_test.go:67:36: BenchmarkSprintfLongString b does not escape
./exconcat_test.go:73:18: BenchmarkSprintfLongString ... argument does not escape
./exconcat_test.go:77:33: BenchmarkPlusLongString b does not escape
./exconcat_test.go:87:33: BenchmarkJoinLongString b does not escape
./exconcat_test.go:93:28: BenchmarkJoinLongString []string literal does not escape
./exconcat_test.go:97:36: BenchmarkBuilderLongString b does not escape
./exconcat_test.go:104:10: BenchmarkBuilderLongString builder does not escape
./exconcat_test.go:105:10: BenchmarkBuilderLongString builder does not escape
./exconcat_test.go:106:14: BenchmarkBuilderLongString builder does not escape
./exconcat_test.go:106:21: BenchmarkBuilderLongString &strings.b·2.buf does not escape
./exconcat_test.go:117:9: buffer escapes to heap
./exconcat_test.go:116:7: moved to heap: buffer
./exconcat_test.go:118:9: buffer escapes to heap
./exconcat_test.go:110:35: BenchmarkBufferLongString b does not escape
./exconcat_test.go:119:13: BenchmarkBufferLongString buffer does not escape
./exconcat_test.go:119:20: BenchmarkBufferLongString string(bytes.b·2.buf[bytes.b·2.off:]) does not escape
./exconcat_test.go:129:18: id escapes to heap
./exconcat_test.go:123:39: BenchmarkSprintfStringInt64_1 b does not escape
./exconcat_test.go:129:18: BenchmarkSprintfStringInt64_1 ... argument does not escape
./exconcat_test.go:139:19: str escapes to heap
./exconcat_test.go:139:19: id escapes to heap
./exconcat_test.go:133:39: BenchmarkSprintfStringInt64_2 b does not escape
./exconcat_test.go:139:18: BenchmarkSprintfStringInt64_2 ... argument does not escape
./exconcat_test.go:143:34: BenchmarkPlusStringInt64 b does not escape
./exconcat_test.go:149:11: BenchmarkPlusStringInt64 str + strconv.FormatInt(id, 10) does not escape
```
