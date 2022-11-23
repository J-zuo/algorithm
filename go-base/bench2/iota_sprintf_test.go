package bench2

import "testing"

func BenchmarkUseStringsItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseStringsItoa()
	}
}

func BenchmarkUseSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseSprintf()
	}
}
