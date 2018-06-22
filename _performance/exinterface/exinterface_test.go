package exinterface

import "testing"

func BenchmarkUser_Scan(b *testing.B) {
	u := New()
	for i := 0; i < b.N; i++ {
		u.Scan("Risal", "Falah")
	}
}

func BenchmarkUser_ScanInterface1(b *testing.B) {
	u := New()
	for i := 0; i < b.N; i++ {
		IPrinter.Scan(&u, "Risal", "Falah")
	}
}

func BenchmarkUser_ScanInterface2(b *testing.B) {
	u := New()
	for i := 0; i < b.N; i++ {
		x := IPrinter(&u)
		x.Scan("Risal", "Falah")
	}
}
