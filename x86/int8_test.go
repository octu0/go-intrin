package x86

import (
	"testing"
)

func BenchmarkInt8(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]int8, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = int8(i % 127)
		}
		tb.ResetTimer()
		_ = int8SumNative(v)
	})
	b.Run("sum/simd", func(tb *testing.B) {
		v := make([]int8, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = int8(i % 127)
		}
		tb.ResetTimer()
		_ = Int8Sum(v...)
	})
}

func BenchmarkUint8(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]uint8, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = uint8(i % 255)
		}
		tb.ResetTimer()
		_ = uint8SumNative(v)
	})
	b.Run("sum/simd", func(tb *testing.B) {
		v := make([]uint8, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = uint8(i % 255)
		}
		tb.ResetTimer()
		_ = Uint8Sum(v...)
	})
}
