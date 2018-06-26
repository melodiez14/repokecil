package exerror

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("Test")
	}
}

func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("Test")
	}
}
