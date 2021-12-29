package main

import (
	"testing"
)

type T = string

var sMakeCopy128 []T

func Benchmark_MakeCopy_128(b *testing.B) {
	s := make([]T, 128)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sMakeCopy128 = make([]T, len(s))
		copy(sMakeCopy128, s)
	}
}

var sAppend128 []T

func Benchmark_Append_128(b *testing.B) {
	s := make([]T, 128)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sAppend128 = append([]T(nil), s...)
	}
	_ = sAppend128
}

/* see benchmark results for the same(MakeCopy vs Append cloning methods - both didn't come out as conclusive winner):

kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               635.2 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               651.1 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.132s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               599.7 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               624.2 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.125s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               623.9 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               632.6 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.128s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               562.7 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               630.0 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.122s

*/
