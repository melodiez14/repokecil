package expool

import (
	"sync"
	"testing"
)

type Astruct struct {
	ID      int
	Name    string
	ID2     int
	Name2   string
	ID3     int
	Name3   string
	ID4     int
	Name4   string
	ID5     int
	Name5   string
	ID6     int
	Name6   string
	ID7     int
	Name7   string
	ID8     int
	Name8   string
	ID9     int
	Name9   string
	ID10    int
	Name10  string
	ID11    int
	Name11  string
	ID12    int
	Name12  string
	_ID     int
	_Name   string
	_ID2    int
	_Name2  string
	_ID3    int
	_Name3  string
	_ID4    int
	_Name4  string
	_ID5    int
	_Name5  string
	_ID6    int
	_Name6  string
	_ID7    int
	_Name7  string
	_ID8    int
	_Name8  string
	_ID9    int
	_Name9  string
	_ID10   int
	_Name10 string
	_ID11   int
	_Name11 string
	_ID12   int
	_Name12 string
}

var pool = &sync.Pool{
	New: func() interface{} {
		return new(Astruct)
	},
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetPool()
	}
}

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Get()
	}
}

//go:noinline
func GetPool() *Astruct {
	x := pool.Get().(*Astruct)
	defer pool.Put(x)
	return x
}

//go:noinline
func Get() *Astruct {
	return &Astruct{}
}
