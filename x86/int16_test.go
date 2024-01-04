package x86

import (
	"testing"
)

func BenchmarkInt16(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]int16, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = int16(i % 32767)
		}
		tb.ResetTimer()
		_ = int16SumNative(v)
	})
	b.Run("sum/simd", func(tb *testing.B) {
		v := make([]int16, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = int16(i % 32767)
		}
		tb.ResetTimer()
		_ = Int16Sum(v...)
	})
}

func BenchmarkUint16(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]uint16, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = uint16(i % 65535)
		}
		tb.ResetTimer()
		_ = uint16SumNative(v)
	})
	b.Run("sum/simd", func(tb *testing.B) {
		v := make([]uint16, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = uint16(i % 65535)
		}
		tb.ResetTimer()
		_ = Uint16Sum(v...)
	})
}
