package exmaps

import "testing"

func BenchmarkSetIntMapsWAllocation(b *testing.B) {
	m := make(map[int]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkSetIntMapsWOAllocation(b *testing.B) {
	m := make(map[int]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkGetIntMapsWAllocation(b *testing.B) {
	m := make(map[int]int, b.N)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := b.N - 1; i >= 0; i-- {
		m[i] = i
	}
}
