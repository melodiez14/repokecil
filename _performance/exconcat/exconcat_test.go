package exconcat

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkSprintfShortString(b *testing.B) {
	const (
		str = "class:room:i_"
		id  = "UDJT206"
	)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", str, id)
	}
}

func BenchmarkPlusShortString(b *testing.B) {
	const (
		str = "class:room:i_"
		id  = "UDJT206"
	)
	for i := 0; i < b.N; i++ {
		_ = str + id
	}
}

func BenchmarkJoinShortString(b *testing.B) {
	const (
		str = "class:room:i_"
		id  = "UDJT206"
	)
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{str, id}, "")
	}
}

func BenchmarkBuilderShortString(b *testing.B) {
	const (
		str = "class:room:i_"
		id  = "UDJT206"
	)
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(str)
		builder.WriteString(id)
		_ = builder.String()
	}
}

func BenchmarkBufferShortString(b *testing.B) {
	const (
		str = "class:room:i_"
		id  = "UDJT206"
	)
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(str)
		buffer.WriteString(id)
		_ = buffer.String()
	}
}

func BenchmarkSprintfLongString(b *testing.B) {
	const (
		strone = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
		strtwo = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", strone, strtwo)
	}
}

func BenchmarkPlusLongString(b *testing.B) {
	const (
		strone = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
		strtwo = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	)
	for i := 0; i < b.N; i++ {
		_ = strone + strtwo
	}
}

func BenchmarkJoinLongString(b *testing.B) {
	const (
		strone = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
		strtwo = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	)
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{strone, strtwo}, "")
	}
}

func BenchmarkBuilderLongString(b *testing.B) {
	const (
		strone = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
		strtwo = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	)
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(strone)
		builder.WriteString(strtwo)
		_ = builder.String()
	}
}

func BenchmarkBufferLongString(b *testing.B) {
	const (
		strone = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
		strtwo = "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	)
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(strone)
		buffer.WriteString(strtwo)
		_ = buffer.String()
	}
}

func BenchmarkSprintfStringInt64_1(b *testing.B) {
	const (
		str = "class:student:i_%d"
		id  = 1408
	)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(str, id)
	}
}

func BenchmarkSprintfStringInt64_2(b *testing.B) {
	const (
		str = "class:student:i_"
		id  = 1408
	)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%d", str, id)
	}
}

func BenchmarkPlusStringInt64(b *testing.B) {
	const (
		str = "class:student:i_"
		id  = 1408
	)
	for i := 0; i < b.N; i++ {
		_ = str + strconv.FormatInt(id, 10)
	}
}
