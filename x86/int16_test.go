package x86

import (
	"reflect"
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

func TestInt16Add(t *testing.T) {
	out := Int16Add(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		2, 5, 32767, -32767, -32767, 32767, -32768, -1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestInt16Sub(t *testing.T) {
	out := Int16Sub(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		0, -1, 32766, -32768, 32767, -32766, -32767, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestInt16Min(t *testing.T) {
	out := Int16Min(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		1, 2, 1, -32768, -32768, 1, -32768, -1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestInt16Max(t *testing.T) {
	out := Int16Max(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		1, 3, 32767, 1, 1, 32767, -1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestInt16MAdd(t *testing.T) {
	out := Int16MAdd(
		[8]int16{1, 2, 3, -1, 32767, -32768, -1, -1},
		[8]int16{1, 5, 2, 0, 1, -1, 0, -1},
	)
	expect := [4]int32{
		11, 6, 65535, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint16Add(t *testing.T) {
	out := Uint16Add(
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 0},
		[8]uint16{1, 3, 1, 1, 65535, 65535, 0, 1},
	)
	expect := [8]uint16{
		2, 5, 32768, 65535, 65535, 65535, 0, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint16Sub(t *testing.T) {
	out := Uint16Sub(
		[8]uint16{0, 1, 1, 0, 65535, 65535, 1, 0},
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 1},
	)
	expect := [8]uint16{
		0, 0, 0, 0, 65534, 0, 1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint16Avg(t *testing.T) {
	out := Uint16Avg(
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 0},
		[8]uint16{1, 3, 1, 1, 65535, 65535, 0, 1},
	)
	expect := [8]uint16{
		1, 3, 16384, 32768, 32768, 65535, 0, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
