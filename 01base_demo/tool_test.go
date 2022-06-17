package base_demo

import "testing"

func TestGetMD5(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		T()
	}
}
func BenchmarkGetMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T()
	}
}

func TestRange(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		randomString()
	}
}
