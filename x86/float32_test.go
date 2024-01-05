package x86

import (
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkFloat32(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]float32, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = float32(i)
		}
		tb.ResetTimer()
		_ = float32SumNative(v)
	})
	b.Run("sum/simd", func(tb *testing.B) {
		v := make([]float32, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = float32(i)
		}
		tb.ResetTimer()
		_ = Float32Sum(v...)
	})
}

func TestFloat32Filter(t *testing.T) {
	t.Run("add", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{128.0, 64.0, 32.0, 16.0}}
		out := f.Add(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			129.0,
			66.0,
			35.0,
			20.0,
			133.0,
			70.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("sub", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{128.0, 64.0, 32.0, 16.0}}
		out := f.Sub(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			-127.0,
			-62.0,
			-29.0,
			-12.0,
			-123.0,
			-58.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("mul", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{2.0, 3.0, 4.0, 8.0}}
		out := f.Mul(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			2.0,
			6.0,
			12.0,
			32.0,
			10.0,
			18.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("div", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{2.0, 3.0, 4.0, 8.0}}
		out := f.Div(
			1.0,
			2.1,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			0.5,
			0.7,
			0.75,
			0.5,
			2.5,
			2.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}

func TestFloat32Add(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := Float32Add(a, b)

	expect := [4]float32{-101.1, 25.2, 0.0, -15.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Add8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.1, 6.2, 7.3, 8.0}
	b := [8]float32{-100.0, 20.0, 0.0, -5.0, 1.0, 2.0, 3.0, 4.0}
	out := Float32Add8(a, b)

	expect := [8]float32{-101.1, 25.2, 0.0, -15.3, 6.1, 8.2, 10.3, 12.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Sub(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := Float32Sub(a, b)

	expect := [4]float32{98.9, -14.8, 0.0, -5.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Sub8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.1, 6.2, 7.3, 8.8}
	b := [8]float32{-100.0, 20.0, 0.0, -5.0, 0.1, 0.2, 0.3, 0.8}
	out := Float32Sub8(a, b)

	expect := [8]float32{98.9, -14.8, 0.0, -5.3, 5.0, 6.0, 7.0, 8.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Mul(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := Float32Mul(a, b)

	expect := [4]float32{110.0, 104.0, 0.0, 51.5}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Mul8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.0, 6.4, 7.0, 8.2}
	b := [8]float32{-100.0, 20.0, 0.0, -5.0, 2.0, 3.0, 4.0, 5.0}
	out := Float32Mul8(a, b)

	expect := [8]float32{110.0, 104.0, 0.0, 51.5, 10.0, 19.2, 28.0, 41.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Div(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.4, -5.0}
	out := Float32Div(a, b)

	expect := [4]float32{0.011, 0.26, 0.0, 2.06}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Div8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.0, 6.3, 7.2, 8.1}
	b := [8]float32{-100.0, 20.0, 0.4, -5.0, 2.0, 2.0, 2.0, 2.0}
	out := Float32Div8(a, b)

	expect := [8]float32{0.011, 0.26, 0.0, 2.06, 2.5, 3.15, 3.6, 4.05}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Sqrt(t *testing.T) {
	a := [4]float32{2.0, 4.0, 13.0, 16.0}
	out := Float32Sqrt(a)

	expect := [4]float32{1.4142135, 2, 3.6055512, 4}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Sqrt8(t *testing.T) {
	a := [8]float32{2.0, 4.0, 13.0, 16.0, 36.0, 64.0, 100.0, 121.0}
	out := Float32Sqrt8(a)

	expect := [8]float32{1.4142135, 2, 3.6055512, 4, 6, 8, 10, 11}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32RSqrt(t *testing.T) {
	a := [4]float32{2.0, 4.0, 13.0, 16.0}
	out := Float32RSqrt(a)

	expect := [4]float32{0.7069092, 0.49987793, 0.2772827, 0.24993896}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32RSqrt8(t *testing.T) {
	a := [8]float32{2.0, 4.0, 13.0, 16.0, 1.2, 2.2, 3.3, 8.2}
	out := Float32RSqrt8(a)

	expect := [8]float32{
		0.7069092, 0.49987793, 0.2772827, 0.24993896,
		0.91296387, 0.67419434, 0.5505371, 0.34924316,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Rcp(t *testing.T) {
	a := [4]float32{1.0, 4.0, 13.0, 100}
	out := Float32Rcp(a)

	expect := [4]float32{0.99975586, 0.24993896, 0.0769043, 0.0099983215}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Rcp8(t *testing.T) {
	a := [8]float32{1.0, 4.0, 13.0, 100, 0.55, 0.6, 0.07, 0.08}
	out := Float32Rcp8(a)

	expect := [8]float32{
		0.99975586, 0.24993896, 0.0769043, 0.0099983215,
		1.8183594, 1.666748, 14.287109, 12.5,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Min(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := Float32Min(a, b)

	expect := [4]float32{-100.0, 5.2, 0.0, -10.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Min8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.2, 6.3, 7.4, 8.5}
	b := [8]float32{-100.0, 20.0, 0.0, -5.0, 5.3, -6.3, 7.0, -1.0}
	out := Float32Min8(a, b)

	expect := [8]float32{-100.0, 5.2, 0.0, -10.3, 5.2, -6.3, 7.0, -1.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Max(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := Float32Max(a, b)

	expect := [4]float32{-1.1, 20.0, 0.0, -5.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Max8(t *testing.T) {
	a := [8]float32{-1.1, 5.2, 0.0, -10.3, 5.2, 6.3, 7.4, 8.5}
	b := [8]float32{-100.0, 20.0, 0.0, -5.0, 5.3, -6.3, 7.0, -1.0}
	out := Float32Max8(a, b)

	expect := [8]float32{-1.1, 20.0, 0.0, -5.0, 5.3, 6.3, 7.4, 8.5}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32And(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
	}
	ub := [4]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_10000001_00000000000000000000000,
		0,
		0b0_10000000_10010000000000000000000,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := Float32And(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32And8(t *testing.T) {
	ua := [8]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
		0b0_10000001_10000000000000000000000,
		0b0_11000001_11000000000000000000000,
		0b0_11100001_11100000000000000000000,
		0b0_11110001_11110000000000000000000,
	}
	ub := [8]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
		0b0_10000000_10000000000000000000000,
		0b0_01000000_10000000000000000000000,
		0b0_01000000_01000000000000000000000,
		0b0_00100000_00100000000000000000000,
	}
	ue := [8]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_10000001_00000000000000000000000,
		0,
		0b0_10000000_10010000000000000000000,
		0b0_10000000_10000000000000000000000,
		0b0_01000000_10000000000000000000000,
		0b0_01000000_01000000000000000000000,
		0b0_00100000_00100000000000000000000,
	}

	a := *(*[8]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[8]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[8]float32)(unsafe.Pointer(&ue[0]))

	out := Float32And8(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32Or(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
	}
	ub := [4]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_11000011_00000000000000000000000,
		0,
		0b0_10000001_11110000000000000000010,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := Float32Or(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32Or8(t *testing.T) {
	ua := [8]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
		0b0_10000001_10000000000000000000000,
		0b0_11000001_11000000000000000000000,
		0b0_11100001_11100000000000000000000,
		0b0_11110001_11110000000000000000000,
	}
	ub := [8]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
		0b0_10000000_10000000000000000000000,
		0b0_01000000_10000000000000000000000,
		0b0_01000000_01000000000000000000000,
		0b0_00100000_00100000000000000000000,
	}
	ue := [8]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_11000011_00000000000000000000000,
		0,
		0b0_10000001_11110000000000000000010,
		0b0_10000001_10000000000000000000000,
		0b0_11000001_11000000000000000000000,
		0b0_11100001_11100000000000000000000,
		0b0_11110001_11110000000000000000000,
	}

	a := *(*[8]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[8]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[8]float32)(unsafe.Pointer(&ue[0]))

	out := Float32Or8(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32Xor(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
	}
	ub := [4]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000010_00000000000000000000000,
		0b0_01000010_00000000000000000000000,
		0,
		0b0_00000001_01100000000000000000010,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := Float32Xor(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32Xor8(t *testing.T) {
	ua := [8]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
		0b0_10000001_11000000000000000000011,
		0b0_10000001_11100000000000000000011,
		0b0_10000001_11110000000000000000011,
		0b0_10000001_11111000000000000000011,
	}
	ub := [8]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
		0b0_10000000_10000000000000000000001,
		0b0_10000000_10000000000000000000001,
		0b0_10000000_10000000000000000000001,
		0b0_10000000_10000000000000000000001,
	}
	ue := [8]uint32{
		0b0_00000010_00000000000000000000000,
		0b0_01000010_00000000000000000000000,
		0,
		0b0_00000001_01100000000000000000010,
		0b0_00000001_01000000000000000000010,
		0b0_00000001_01100000000000000000010,
		0b0_00000001_01110000000000000000010,
		0b0_00000001_01111000000000000000010,
	}

	a := *(*[8]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[8]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[8]float32)(unsafe.Pointer(&ue[0]))

	out := Float32Xor8(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32AndNot(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
	}
	ub := [4]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000010_00000000000000000000000,
		0b0_00000010_00000000000000000000000,
		0,
		0b0_00000000_01100000000000000000000,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := Float32AndNot(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32AndNot8(t *testing.T) {
	ua := [8]uint32{
		0b0_00000001_00000000000000000000000,
		0b0_11000001_00000000000000000000000,
		0,
		0b0_10000001_10010000000000000000010,
		0b0_00000001_00000000000000000000000,
		0b0_00000001_00000000000000000000000,
		0b0_00000001_00000000000000000000000,
		0b0_00000001_00000000000000000000000,
	}
	ub := [8]uint32{
		0b0_00000011_00000000000000000000000,
		0b0_10000011_00000000000000000000000,
		0,
		0b0_10000000_11110000000000000000000,
		0b0_00000011_10000000000000000000000,
		0b0_00000011_11000000000000000000000,
		0b0_00000011_11100000000000000000000,
		0b0_00000011_01000000000000000000000,
	}
	ue := [8]uint32{
		0b0_00000010_00000000000000000000000,
		0b0_00000010_00000000000000000000000,
		0,
		0b0_00000000_01100000000000000000000,
		0b0_00000010_10000000000000000000000,
		0b0_00000010_11000000000000000000000,
		0b0_00000010_11100000000000000000000,
		0b0_00000010_01000000000000000000000,
	}

	a := *(*[8]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[8]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[8]float32)(unsafe.Pointer(&ue[0]))

	out := Float32AndNot8(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat32Sum(t *testing.T) {
	t.Run("size=4", func(tt *testing.T) {
		out := Float32Sum(1.0, 2.0, 3.0, 4.0)
		expect := float32(10.0)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("odd", func(tt *testing.T) {
		out := Float32Sum(1.0, 2.0, 3.0, 4.0, 5.0)
		expect := float32(15.0)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("large", func(tt *testing.T) {
		N := 1023
		v := make([]float32, N)
		for i := 0; i < N; i += 1 {
			v[i] = float32(i)
		}

		out := Float32Sum(v...)
		expect := float32SumNative(v)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}
