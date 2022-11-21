package channel

import "testing"

//基准测试使用
func BenchmarkUseMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseMutex()
	}
}

func BenchmarkUseRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseRWMutex()
	}
}

func BenchmarkUseChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseChan()
	}
}
