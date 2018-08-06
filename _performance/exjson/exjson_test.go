package main

import (
	"encoding/json"
	"testing"

	"github.com/valyala/quicktemplate"

	"github.com/json-iterator/go"
	"github.com/melodiez14/repokecil/_performance/exjson/templates"
)

type SmallStruct struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int8   `json:"age"`
}

func BenchmarkSmallJSONIter(b *testing.B) {
	var a = SmallStruct{
		Age:       10,
		FirstName: "Risal",
		LastName:  "Falah",
	}
	cfg := jsoniter.Config{}.Froze()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg.Marshal(&a)
	}
}

func BenchmarkSmallStd(b *testing.B) {
	var a = SmallStruct{
		Age:       10,
		FirstName: "Risal",
		LastName:  "Falah",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(&a)
	}
}

func BenchmarkLargeJSONIter(b *testing.B) {
	var a = []SmallStruct{
		SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		},
	}
	cfg := jsoniter.Config{}.Froze()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg.Marshal(&a)
	}
}

func BenchmarkLargeStd(b *testing.B) {
	var a = []SmallStruct{
		SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		},
	}
	cfg := jsoniter.ConfigCompatibleWithStandardLibrary
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg.Marshal(&a)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		templates.Hello("Risal")
	}
}

func BenchmarkSmallQuick(b *testing.B) {
	var a = templates.SmallStruct{
		Age:       10,
		FirstName: "Risal",
		LastName:  "Falah",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := quicktemplate.AcquireByteBuffer()
		a.WriteJSON(x)
		quicktemplate.ReleaseByteBuffer(x)
	}
}

func BenchmarkLargeQuick(b *testing.B) {
	var a = templates.LargeStruct{
		templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		}, templates.SmallStruct{
			Age:       10,
			FirstName: "Risal",
			LastName:  "Falah",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.JSON()
	}
}
