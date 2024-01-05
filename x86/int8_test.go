package x86

import (
	"reflect"
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

func TestInt8Add(t *testing.T) {
	out := Int8Add(
		[16]int8{
			1, 2, 3, 4, 5, 6, 7, 8,
			-1, 2, 3, -4, -5, -6, -7, -8,
		},
		[16]int8{
			11, 12, 13, 14, 15, 16, 17, 18,
			100, 125, 126, -128, -124, 127, -2, -3,
		},
	)
	expect := [16]int8{
		12, 14, 16, 18, 20, 22, 24, 26,
		99, 127, 127, -128, -128, 121, -9, -11,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestInt8Sub(t *testing.T) {
	out := Int8Sub(
		[16]int8{
			1, 2, 3, 4, 5, 6, 7, 8,
			127, -128, -3, -127, 0, -1, 1, 8,
		},
		[16]int8{
			1, -2, 1, -2, 0, -1, 1, -8,
			-128, -1, 127, -127, -128, 1, -1, 8,
		},
	)
	expect := [16]int8{
		0, 4, 2, 6, 5, 7, 6, 16,
		127, -127, -128, 0, 127, -2, 2, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8Add(t *testing.T) {
	out := Uint8Add(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			1, 2, 3, 4, 5, 6, 7, 8,
		},
		[16]uint8{
			1, 2, 3, 4, 0, 0, 7, 8,
			100, 254, 255, 251, 0, 1, 127, 128,
		},
	)
	expect := [16]uint8{
		2, 4, 6, 8, 5, 6, 14, 16,
		101, 255, 255, 255, 5, 7, 134, 136,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8Sub(t *testing.T) {
	out := Uint8Sub(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 254, 255, 255, 127, 6, 1, 8,
		},
		[16]uint8{
			1, 0, 255, 0, 4, 0, 8, 8,
			1, 2, 255, 0, 127, 6, 0, 8,
		},
	)
	expect := [16]uint8{
		0, 2, 0, 4, 1, 6, 0, 0,
		254, 252, 0, 255, 0, 0, 1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8Avg(t *testing.T) {
	out := Uint8Avg(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 255, 255, 127, 6, 1, 8,
		},
		[16]uint8{
			9, 2, 4, 16, 127, 121, 121, 16,
			255, 5, 0, 127, 127, 12, 254, 255,
		},
	)
	expect := [16]uint8{
		5, 2, 4, 10, 66, 64, 64, 12,
		255, 128, 128, 191, 127, 9, 128, 132,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8Min(t *testing.T) {
	out := Uint8Min(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 0, 255, 127, 6, 7, 8,
		},
		[16]uint8{
			0, 1, 10, 255, 254, 0, 7, 8,
			255, 251, 255, 1, 128, 7, 6, 9,
		},
	)
	expect := [16]uint8{
		0, 1, 3, 4, 5, 0, 7, 8,
		255, 250, 0, 1, 127, 6, 6, 8,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8Max(t *testing.T) {
	out := Uint8Max(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 0, 255, 127, 6, 7, 8,
		},
		[16]uint8{
			0, 1, 10, 255, 254, 0, 7, 8,
			255, 251, 255, 1, 128, 7, 6, 9,
		},
	)
	expect := [16]uint8{
		1, 2, 10, 255, 254, 6, 7, 8,
		255, 251, 255, 255, 128, 7, 7, 9,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
