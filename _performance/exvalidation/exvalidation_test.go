package exvalidation

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/go-playground/form"
)

type BStruct struct {
	Age string
}

type AStruct struct {
	Age int64 `form:"age"`
}

type validator interface {
	Validate()
}

func BenchmarkFormValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "1234"
		integer, _ := strconv.ParseInt(str, 10, 64)
		_ = AStruct{Age: integer}
	}
}

func BenchmarkValidator(b *testing.B) {
	values := url.Values{
		"age": []string{"1234"},
	}
	decoder := form.NewDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var astruct AStruct
		decoder.Decode(&astruct, values)
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bstruct := BStruct{Age: "1234"}
		integer, _ := strconv.ParseInt(bstruct.Age, 10, 64)
		_ = AStruct{Age: integer}
	}
}
