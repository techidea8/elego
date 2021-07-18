package utils

import (
	"testing"
)

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(32)
	}
}
