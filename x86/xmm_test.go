package x86

import (
	"testing"
)

func BenchmarkXmm(b *testing.B) {
	b.Run("cgo", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			_ = xmmAdd([4]float32{}, [4]float32{})
		}
	})
	b.Run("link", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			_ = __xmmAdd([4]float32{}, [4]float32{})
		}
	})
}

func TestXmm(t *testing.T) {
	t.Run("cgo", func(tt *testing.T) {
		out := xmmAdd(
			[4]float32{1.0, 2.0, 3.0, 4.0},
			[4]float32{2.0, 3.0, 4.0, 5.0},
		)
		if (out[0] == 3.0 &&
			out[1] == 5.0 &&
			out[2] == 7.0 &&
			out[3] == 9.0) != true {
			tt.Errorf("actual %v", out)
		}
	})
	t.Run("link", func(tt *testing.T) {
		out := __xmmAdd(
			[4]float32{1.0, 2.0, 3.0, 4.0},
			[4]float32{2.0, 3.0, 4.0, 5.0},
		)
		if (out[0] == 3.0 &&
			out[1] == 5.0 &&
			out[2] == 7.0 &&
			out[3] == 9.0) != true {
			tt.Errorf("actual %v", out)
		}
	})
}
