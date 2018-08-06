package exhash

import (
	"testing"

	"github.com/cespare/xxhash"

	"golang.org/x/crypto/sha3"
)

func BenchmarkSHA3_224(b *testing.B) {
	hash := sha3.New224()
	payload := []byte(`{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}`)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		digest := [28]byte{}
		hash.Write(payload)
		hash.Sum(digest[:0])
		hash.Reset()
	}
}

func BenchmarkXXHash(b *testing.B) {
	str := `{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}{"user_id: 12345"}`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xxhash.Sum64String(str)
	}
}
